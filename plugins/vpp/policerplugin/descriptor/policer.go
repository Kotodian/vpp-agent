package descriptor

import (
	"fmt"

	"go.ligato.io/cn-infra/v2/idxmap"
	"go.ligato.io/cn-infra/v2/logging"
	"go.ligato.io/vpp-agent/v3/pkg/models"
	kvs "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/descriptor/adapter"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/policeridx"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/vppcalls"
	policer "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/policer"
	"google.golang.org/protobuf/proto"
)

const (
	// PolicerDescriptorName is the name of the descriptor for VPP policer.
	PolicerDescriptorName = "vpp-policer"
)

// PolicerDescriptor teaches KVScheduler how to configure VPP policer.
type PolicerDescriptor struct {
	log            logging.Logger
	policerHandler vppcalls.PolicerVppAPI
}

// NewPolicerDescriptor creates a new instance of the policer descriptor.
func NewPolicerDescriptor(policerHandler vppcalls.PolicerVppAPI, log logging.PluginLogger) *PolicerDescriptor {
	return &PolicerDescriptor{
		policerHandler: policerHandler,
		log:            log.NewLogger("policer-descriptor"),
	}
}

// GetDescriptor returns descriptor suitable for registration (via adapter) with
// the KVScheduler.
func (d *PolicerDescriptor) GetDescriptor() *adapter.PolicerDescriptor {
	return &adapter.PolicerDescriptor{
		Name:          PolicerDescriptorName,
		NBKeyPrefix:   policer.ModelPolicerConfig.KeyPrefix(),
		ValueTypeName: policer.ModelPolicerConfig.ProtoName(),
		// KeyLabel as metadata map key
		KeyLabel:        policer.ModelPolicerConfig.StripKeyPrefix,
		KeySelector:     policer.ModelPolicerConfig.IsKeyValid,
		ValueComparator: d.EquivalentPolicers,
		Validate:        d.Validate,
		DerivedValues:   d.DerivedValues,
		Create:          d.Create,
		Update:          d.Update,
		Delete:          d.Delete,
		Retrieve:        d.Retrieve,
		WithMetadata:    true,
		MetadataMapFactory: func() idxmap.NamedMappingRW {
			return policeridx.NewPolicerIndex(d.log, "vpp-policer-index")
		},
	}
}

func (d *PolicerDescriptor) EquivalentPolicers(key string, oldPolicer, newPolicer *policer.PolicerConfig) bool {
	// compare base fields
	return proto.Equal(oldPolicer, newPolicer)
}

func (d *PolicerDescriptor) Validate(key string, policer *policer.PolicerConfig) (err error) {
	//TODO: add validation
	return nil
}

// DerivedValues derives policer.PolicerConfig_Interface for every interface assigned to the Policer.
func (d *PolicerDescriptor) DerivedValues(key string, p *policer.PolicerConfig) (derValues []kvs.KeyValuePair) {
	// Policer interfaces
	for _, policerIface := range p.Interfaces {
		derValues = append(derValues, kvs.KeyValuePair{
			Key:   policer.DerivedPolicerInterfaceKey(p.Name, policerIface.Name, policerIface.IsOutput),
			Value: policerIface,
		})
	}

	return derValues
}

// Create adds a new policer.
func (d *PolicerDescriptor) Create(key string, policer *policer.PolicerConfig) (metadata *policeridx.PolicerMetadata, err error) {
	var vppPolicerIndex uint32
	vppPolicerIndex, err = d.policerHandler.AddPolicer(policer)
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	metadata = &policeridx.PolicerMetadata{
		Index: vppPolicerIndex,
	}

	return metadata, err
}

// Update updates a existing policer.
func (d *PolicerDescriptor) Update(key string, oldPolicer *policer.PolicerConfig, newPolicer *policer.PolicerConfig, oldMetadata *policeridx.PolicerMetadata) (metadata *policeridx.PolicerMetadata, err error) {
	if oldMetadata == nil {
		return nil, fmt.Errorf("failed to update policer - metadata is nil")
	}

	err = d.policerHandler.UpdatePolicer(oldMetadata.Index, newPolicer)
	if err != nil {
		return nil, err
	}

	return oldMetadata, nil
}

// Delete removes a policer.
func (d *PolicerDescriptor) Delete(key string, policer *policer.PolicerConfig, metadata *policeridx.PolicerMetadata) error {
	if metadata == nil {
		return fmt.Errorf("failed to delete policer - metadata is nil")
	}

	err := d.policerHandler.DelPolicer(metadata.Index)
	if err != nil {
		d.log.Error(err)
	}
	return err
}

// Retrieve returns all wg peers.
func (d *PolicerDescriptor) Retrieve(correlate []adapter.PolicerKVWithMetadata) (dump []adapter.PolicerKVWithMetadata, err error) {
	policers, err := d.policerHandler.DumpPolicers()
	if err != nil {
		d.log.Error(err)
		return dump, err
	}
	for _, policer := range policers {
		dump = append(dump, adapter.PolicerKVWithMetadata{
			Metadata: policer.Metadata,
			Key:      models.Key(policer.Config),
			Value:    policer.Config,
			Origin:   kvs.FromNB,
		})
	}

	return dump, nil
}

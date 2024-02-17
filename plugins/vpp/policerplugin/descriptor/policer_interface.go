package descriptor

import (
	"github.com/pkg/errors"
	"go.ligato.io/cn-infra/v2/logging"
	kvs "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin/ifaceidx"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/descriptor/adapter"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/policeridx"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/vppcalls"
	interfaces "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/interfaces"
	policer "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/policer"
	"google.golang.org/protobuf/proto"
)

const (
	// PolicerInterfaceDescriptorName is the name of the descriptor for VPP policer.
	PolicerInterfaceDescriptorName = "vpp-policer-interface"
	// dependency labels
	policerInterfaceDep = "interface-exists"
)

// PolicerInterfaceDescriptor teaches KVScheduler how to configure VPP policer interface.
type PolicerInterfaceDescriptor struct {
	log            logging.Logger
	policerHandler vppcalls.PolicerVppAPI
	policerIndex   policeridx.PolicerMetadataIndex
	ifIndex        ifaceidx.IfaceMetadataIndex
}

// NewPolicerInterfaceDescriptor creates a new instance of the policer descriptor.
func NewPolicerInterfaceDescriptor(ifIndex ifaceidx.IfaceMetadataIndex, policerIndex policeridx.PolicerMetadataIndex, policerHandler vppcalls.PolicerVppAPI, log logging.PluginLogger) *PolicerInterfaceDescriptor {
	return &PolicerInterfaceDescriptor{
		log:            log.NewLogger("policer-interface-descriptor"),
		ifIndex:        ifIndex,
		policerIndex:   policerIndex,
		policerHandler: policerHandler,
	}
}

func (d *PolicerInterfaceDescriptor) GetDescriptor() *adapter.PolicerInterfaceDescriptor {
	return &adapter.PolicerInterfaceDescriptor{
		Name: PolicerInterfaceDescriptorName,
		KeySelector: func(key string) bool {
			_, _, _, isPolicerInterfaceKey := policer.ParseDerivedPolicerInterfaceKey(key)
			return isPolicerInterfaceKey
		},
		ValueTypeName: string(proto.MessageName(&policer.PolicerConfig{})),
		Create:        d.Create,
		Delete:        d.Delete,
		Dependencies:  d.Dependencies,
	}
}

// Create puts interface into policer.
func (d *PolicerInterfaceDescriptor) Create(key string, policerIf *policer.PolicerConfig_Interface) (metadata interface{}, err error) {
	// get policer name
	policerName, ifName, isOutput, isPolicerInterfaceKey := policer.ParseDerivedPolicerInterfaceKey(key)
	if !isPolicerInterfaceKey {
		err = errors.Errorf("provided key is not a derived Policer <=> interface binding key %s", key)
		d.log.Error(err)
		return nil, err
	}

	ifMetadata, exists := d.ifIndex.LookupByName(ifName)
	if !exists {
		err = errors.Errorf("interface name: %s doesn't exists", ifName)
		d.log.Error(err)
		return nil, err
	}

	policerMetadata, exists := d.policerIndex.LookupByName(policerName)
	if !exists {
		err = errors.Errorf("policer name: %s doesn't exists", policerName)
		d.log.Error(err)
		return nil, err
	}

	// put interface into the policer
	if isOutput {
		err = d.policerHandler.PolicerOutput(ifMetadata.SwIfIndex, policerMetadata.GetIndex(), true)
		if err != nil {
			d.log.Error(err)
			return nil, err
		}
	} else {
		err = d.policerHandler.PolicerInput(ifMetadata.SwIfIndex, policerMetadata.GetIndex(), true)
		if err != nil {
			d.log.Error(err)
			return nil, err
		}
	}

	return nil, nil
}

// Delete removes interface into policer.
func (d *PolicerInterfaceDescriptor) Delete(key string, policerIf *policer.PolicerConfig_Interface, _ interface{}) (err error) {
	// get policer name
	policerName, ifName, isOutput, isPolicerInterfaceKey := policer.ParseDerivedPolicerInterfaceKey(key)
	if !isPolicerInterfaceKey {
		err = errors.Errorf("provided key is not a derived Policer <=> interface binding key %s", key)
		d.log.Error(err)
		return err
	}

	ifMetadata, exists := d.ifIndex.LookupByName(ifName)
	if !exists {
		err = errors.Errorf("interface name: %s doesn't exists", ifName)
		d.log.Error(err)
		return err
	}

	policerMetadata, exists := d.policerIndex.LookupByName(policerName)
	if !exists {
		err = errors.Errorf("policer name: %s doesn't exists", policerName)
		d.log.Error(err)
		return err
	}

	// put interface into the policer
	if isOutput {
		err = d.policerHandler.PolicerOutput(ifMetadata.SwIfIndex, policerMetadata.GetIndex(), false)
		if err != nil {
			d.log.Error(err)
			return err
		}
	} else {
		err = d.policerHandler.PolicerInput(ifMetadata.SwIfIndex, policerMetadata.GetIndex(), false)
		if err != nil {
			d.log.Error(err)
			return err
		}
	}

	return nil
}

// Dependencies lists the interface as the only dependency for the binding.
func (d *PolicerInterfaceDescriptor) Dependencies(key string, value *policer.PolicerConfig_Interface) []kvs.Dependency {
	return []kvs.Dependency{
		{
			Label: policerInterfaceDep,
			Key:   interfaces.InterfaceKey(value.Name),
		},
	}
}

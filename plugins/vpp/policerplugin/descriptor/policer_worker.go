package descriptor

import (
	"github.com/pkg/errors"
	"go.ligato.io/cn-infra/v2/logging"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/descriptor/adapter"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/policeridx"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/vppcalls"
	policer "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/policer"
	"google.golang.org/protobuf/proto"
)

const (
	// PolicerWorkerDescriptorName is the name of the descriptor for VPP policer.
	PolicerWorkerDescriptorName = "vpp-policer-worker"
)

type PolicerWorkerDescriptor struct {
	log            logging.Logger
	policerHandler vppcalls.PolicerVppAPI
	policerIndex   policeridx.PolicerMetadataIndex
}

// NewPolicerWorkerDescriptor creates a new instance of the policer descriptor.
func NewPolicerWorkerDescriptor(policerIndex policeridx.PolicerMetadataIndex, policerHandler vppcalls.PolicerVppAPI, log logging.PluginLogger) *PolicerWorkerDescriptor {
	return &PolicerWorkerDescriptor{
		log:            log.NewLogger("policer-worker-descriptor"),
		policerIndex:   policerIndex,
		policerHandler: policerHandler,
	}
}

func (d *PolicerWorkerDescriptor) GetDescriptor() *adapter.PolicerWorkerDescriptor {
	return &adapter.PolicerWorkerDescriptor{
		Name: PolicerWorkerDescriptorName,
		KeySelector: func(key string) bool {
			_, _, isPolicerWorkerKey := policer.ParseDerivedPolicerWorkerKey(key)
			return isPolicerWorkerKey
		},
		ValueTypeName: string(proto.MessageName(&policer.PolicerConfig{})),
		Create:        d.Create,
		Delete:        d.Delete,
	}
}

// Create puts policer into worker.
func (d *PolicerWorkerDescriptor) Create(key string, policerWorker *policer.PolicerConfig_Worker) (metadata interface{}, err error) {
	// get policer name
	policerName, _, isPolicerWorkerKey := policer.ParseDerivedPolicerWorkerKey(key)
	if !isPolicerWorkerKey {
		err = errors.Errorf("provided key is not a derived Policer <=> worker binding key %s", key)
		d.log.Error(err)
		return nil, err
	}

	policerMetadata, exists := d.policerIndex.LookupByName(policerName)
	if !exists {
		err = errors.Errorf("policer name: %s doesn't exists", policerName)
		d.log.Error(err)
		return nil, err
	}

	err = d.policerHandler.PolicerBind(policerMetadata.GetIndex(), policerWorker, true)
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	return nil, nil
}

// Delete puts policer out of worker.
func (d *PolicerWorkerDescriptor) Delete(key string, policerWorker *policer.PolicerConfig_Worker, _ interface{}) (err error) {
	// get policer name
	policerName, _, isPolicerWorkerKey := policer.ParseDerivedPolicerWorkerKey(key)
	if !isPolicerWorkerKey {
		err = errors.Errorf("provided key is not a derived Policer <=> worker binding key %s", key)
		d.log.Error(err)
		return err
	}

	policerMetadata, exists := d.policerIndex.LookupByName(policerName)
	if !exists {
		err = errors.Errorf("policer name: %s doesn't exists", policerName)
		d.log.Error(err)
		return err
	}

	err = d.policerHandler.PolicerBind(policerMetadata.GetIndex(), policerWorker, false)
	if err != nil {
		d.log.Error(err)
		return err
	}

	return nil
}

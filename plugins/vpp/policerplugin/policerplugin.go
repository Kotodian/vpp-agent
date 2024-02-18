//go:generate descriptor-adapter --descriptor-name Policer --value-type *vpp_policer.PolicerConfig --meta-type *policeridx.PolicerMetadata --import "go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/policeridx" --import "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/policer" --output-dir "descriptor"
//go:generate descriptor-adapter --descriptor-name PolicerInterface --value-type *vpp_policer.PolicerConfig_Interface --import "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/policer" --output-dir "descriptor"

package policerplugin

import (
	"errors"

	"go.ligato.io/cn-infra/v2/health/statuscheck"
	"go.ligato.io/cn-infra/v2/infra"
	"go.ligato.io/vpp-agent/v3/plugins/govppmux"
	kvs "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"

	"go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/descriptor"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/descriptor/adapter"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/policeridx"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/vppcalls"

	_ "go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/vppcalls/vpp2306"
	_ "go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/vppcalls/vpp2310"
)

var DefaultPlugin = *NewPlugin()

type PolicerPlugin struct {
	Deps
	// handler
	PolicerHandler vppcalls.PolicerVppAPI

	policerDescriptor   *descriptor.PolicerDescriptor
	policerIfDescriptor *descriptor.PolicerInterfaceDescriptor

	// runime
	policerIndex policeridx.PolicerMetadataIndex
}

type Deps struct {
	infra.PluginDeps
	KVScheduler kvs.KVScheduler
	VPP         govppmux.API
	IfPlugin    ifplugin.API
	StatusCheck statuscheck.PluginStatusWriter // optional
}

func (p *PolicerPlugin) Init() (err error) {
	// init policer handler
	p.PolicerHandler = vppcalls.CompatiblePolicerVppHandler(p.VPP, p.IfPlugin.GetInterfaceIndex(), p.Log)
	if p.PolicerHandler == nil {
		return errors.New("Policer handler is not available")
	}

	p.policerDescriptor = descriptor.NewPolicerDescriptor(p.PolicerHandler, p.Log)
	policerDescriptor := adapter.NewPolicerDescriptor(p.policerDescriptor.GetDescriptor())
	err = p.KVScheduler.RegisterKVDescriptor(policerDescriptor)
	if err != nil {
		return err
	}

	metadataMap := p.KVScheduler.GetMetadataMap(p.policerDescriptor.GetDescriptor().Name)

	var withIndex bool
	p.policerIndex, withIndex = metadataMap.(policeridx.PolicerMetadataIndex)
	if !withIndex {
		return errors.New("missing index with policer metadata")
	}
	p.policerIfDescriptor = descriptor.NewPolicerInterfaceDescriptor(p.GetPolicerIndex(), p.PolicerHandler, p.Log)
	policerIfDescriptor := adapter.NewPolicerInterfaceDescriptor(p.policerIfDescriptor.GetDescriptor())
	err = p.KVScheduler.RegisterKVDescriptor(policerIfDescriptor)
	if err != nil {
		return err
	}
	return nil
}

func (p *PolicerPlugin) GetPolicerIndex() policeridx.PolicerMetadataIndex {
	return p.policerIndex
}

// AfterInit
func (p *PolicerPlugin) AfterInit() error {
	if p.StatusCheck != nil {
		p.StatusCheck.Register(p.PluginName, nil)
	}
	return nil
}

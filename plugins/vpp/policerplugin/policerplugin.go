//go:generate descriptor-adapter --descriptor-name Policer --value-type *vpp_policer.PolicerConfig --meta-type *policeridx.PolicerMetadata --import "go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/policeridx" --import "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/policer" --output-dir "descriptor"

package policerplugin

import (
	"errors"

	"go.ligato.io/cn-infra/v2/infra"
	"go.ligato.io/vpp-agent/v3/plugins/govppmux"
	kvs "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"

	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/descriptor"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/descriptor/adapter"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/vppcalls"
	_ "go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/vppcalls/vpp2310"
)

var DefaultPlugin = *NewPlugin()

type PolicerPlugin struct {
	Deps
	// handler
	PolicerHandler vppcalls.PolicerVppAPI

	policerDescriptor *descriptor.PolicerDescriptor
}

type Deps struct {
	infra.PluginDeps
	KVScheduler kvs.KVScheduler
	VPP         govppmux.API
}

func (p *PolicerPlugin) Init() (err error) {
	// init policer handler
	p.PolicerHandler = vppcalls.CompatiblePolicerVppHandler(p.VPP, p.Log)
	if p.PolicerHandler == nil {
		return errors.New("Policer handler is not available")
	}

	p.policerDescriptor = descriptor.NewPolicerDescriptor(p.PolicerHandler, p.Log)
	policerDescriptor := adapter.NewPolicerDescriptor(p.policerDescriptor.GetDescriptor())
	err = p.KVScheduler.RegisterKVDescriptor(policerDescriptor)
	if err != nil {
		return err
	}

	return nil
}

// AfterInit
func (p *PolicerPlugin) AfterInit() error {
	return nil
}

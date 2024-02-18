package policerplugin

import (
	"go.ligato.io/cn-infra/v2/health/statuscheck"
	"go.ligato.io/cn-infra/v2/logging"
	"go.ligato.io/vpp-agent/v3/plugins/govppmux"
	"go.ligato.io/vpp-agent/v3/plugins/kvscheduler"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin"
)

// Option is a function that can be used in NewPlugin to customize Plugin.
type Option func(*PolicerPlugin)

// NewPlugin creates a new Plugin with the provided Options.
func NewPlugin(opts ...Option) *PolicerPlugin {
	p := &PolicerPlugin{}

	p.PluginName = "vpp-policer-plugin"
	p.KVScheduler = &kvscheduler.DefaultPlugin
	p.VPP = &govppmux.DefaultPlugin
	p.IfPlugin = &ifplugin.DefaultPlugin
	p.StatusCheck = &statuscheck.DefaultPlugin

	for _, o := range opts {
		o(p)
	}

	if p.Log == nil {
		p.Log = logging.ForPlugin(p.String())
	}

	return p
}

// UseDeps returns Option that can inject custom dependencies.
func UseDeps(f func(*Deps)) Option {
	return func(p *PolicerPlugin) {
		f(&p.Deps)
	}
}

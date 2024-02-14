package vppcalls

import (
	govppapi "go.fd.io/govpp/api"
	"go.ligato.io/cn-infra/v2/logging"
	"go.ligato.io/vpp-agent/v3/plugins/vpp"
	policer "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/policer"
)

type PolicerVppAPI interface {
	PolicerVppRead

	// Add policer via binary API
	AddPolicer(cfg *policer.PolicerConfig) (uint32, error)
	// Update policer via binary API
	UpdatePolicer(policer_index uint32, cfg *policer.PolicerConfig) error
	// Del policer via binary API
	DelPolicer(policer_index uint32) error
}

// PolicerVPPRead provides read methods for policer
type PolicerVppRead interface {
	DumpPolicers() (policerList []*policer.PolicerConfig, err error)
}

var Handler = vpp.RegisterHandler(vpp.HandlerDesc{
	Name:       "policer",
	HandlerAPI: (*PolicerVppAPI)(nil),
})

type NewHandlerFunc func(ch govppapi.Channel, log logging.Logger) PolicerVppAPI

func AddHandlerVersion(version vpp.Version, msgs []govppapi.Message, h NewHandlerFunc) {
	Handler.AddVersion(vpp.HandlerVersion{
		Version: version,
		Check: func(c vpp.Client) error {
			ch, err := c.NewAPIChannel()
			if err != nil {
				return err
			}
			return ch.CheckCompatiblity(msgs...)
		},
		NewHandler: func(c vpp.Client, a ...interface{}) vpp.HandlerAPI {
			ch, err := c.NewAPIChannel()
			if err != nil {
				return err
			}
			return h(ch, a[0].(logging.Logger))
		},
	})
}

func CompatiblePolicerVppHandler(c vpp.Client, log logging.Logger) PolicerVppAPI {
	if v := Handler.FindCompatibleVersion(c); v != nil {
		return v.NewHandler(c, log).(PolicerVppAPI)
	}
	return nil
}

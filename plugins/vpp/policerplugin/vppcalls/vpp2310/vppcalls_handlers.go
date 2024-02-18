package vpp2310

import (
	govppapi "go.fd.io/govpp/api"
	"go.ligato.io/cn-infra/v2/logging"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310"
	vpp_policer "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/policer"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin/ifaceidx"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/vppcalls"
)

func init() {
	var msgs []govppapi.Message

	msgs = append(msgs, vpp_policer.AllMessages()...)

	vppcalls.AddHandlerVersion(vpp2310.Version, msgs, NewPolicerVppHandler)
}

type PolicerVppHandler struct {
	callsChannel govppapi.Channel
	ifIdx        ifaceidx.IfaceMetadataIndex
	log          logging.Logger
}

func NewPolicerVppHandler(ch govppapi.Channel, ifIdx ifaceidx.IfaceMetadataIndex, log logging.Logger) vppcalls.PolicerVppAPI {
	return &PolicerVppHandler{ch, ifIdx, log}
}

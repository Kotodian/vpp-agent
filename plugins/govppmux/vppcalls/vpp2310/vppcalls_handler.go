package vpp2310

import (
	"go.ligato.io/vpp-agent/v3/plugins/govppmux/vppcalls"
	"go.ligato.io/vpp-agent/v3/plugins/vpp"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/memclnt"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/vlib"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/vpe"
)

func init() {
	msgs := vpp.Messages(
		vpe.AllMessages,
		memclnt.AllMessages,
		vlib.AllMessages,
	)
	vppcalls.AddVersion(vpp2310.Version, msgs.AllMessages(), NewVpeHandler)
}

type VpeHandler struct {
	memclnt memclnt.RPCService
	vlib    vlib.RPCService
	vpe     vpe.RPCService
}

func NewVpeHandler(c vpp.Client) vppcalls.VppCoreAPI {
	return &VpeHandler{
		memclnt: memclnt.NewServiceClient(c),
		vlib:    vlib.NewServiceClient(c),
		vpe:     vpe.NewServiceClient(c),
	}
}

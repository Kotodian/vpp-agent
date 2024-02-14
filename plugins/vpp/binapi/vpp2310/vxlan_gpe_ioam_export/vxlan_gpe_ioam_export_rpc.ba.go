// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package vxlan_gpe_ioam_export

import (
	"context"

	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service vxlan_gpe_ioam_export.
type RPCService interface {
	VxlanGpeIoamExportEnableDisable(ctx context.Context, in *VxlanGpeIoamExportEnableDisable) (*VxlanGpeIoamExportEnableDisableReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) VxlanGpeIoamExportEnableDisable(ctx context.Context, in *VxlanGpeIoamExportEnableDisable) (*VxlanGpeIoamExportEnableDisableReply, error) {
	out := new(VxlanGpeIoamExportEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

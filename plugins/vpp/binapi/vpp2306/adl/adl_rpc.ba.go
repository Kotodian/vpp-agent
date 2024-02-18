// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package adl

import (
	"context"

	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service adl.
type RPCService interface {
	AdlAllowlistEnableDisable(ctx context.Context, in *AdlAllowlistEnableDisable) (*AdlAllowlistEnableDisableReply, error)
	AdlInterfaceEnableDisable(ctx context.Context, in *AdlInterfaceEnableDisable) (*AdlInterfaceEnableDisableReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) AdlAllowlistEnableDisable(ctx context.Context, in *AdlAllowlistEnableDisable) (*AdlAllowlistEnableDisableReply, error) {
	out := new(AdlAllowlistEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) AdlInterfaceEnableDisable(ctx context.Context, in *AdlInterfaceEnableDisable) (*AdlInterfaceEnableDisableReply, error) {
	out := new(AdlInterfaceEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
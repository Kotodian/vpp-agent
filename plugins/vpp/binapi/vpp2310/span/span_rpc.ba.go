// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package span

import (
	"context"
	"fmt"
	"io"

	api "go.fd.io/govpp/api"
	memclnt "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/memclnt"
)

// RPCService defines RPC service span.
type RPCService interface {
	SwInterfaceSpanDump(ctx context.Context, in *SwInterfaceSpanDump) (RPCService_SwInterfaceSpanDumpClient, error)
	SwInterfaceSpanEnableDisable(ctx context.Context, in *SwInterfaceSpanEnableDisable) (*SwInterfaceSpanEnableDisableReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) SwInterfaceSpanDump(ctx context.Context, in *SwInterfaceSpanDump) (RPCService_SwInterfaceSpanDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SwInterfaceSpanDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SwInterfaceSpanDumpClient interface {
	Recv() (*SwInterfaceSpanDetails, error)
	api.Stream
}

type serviceClient_SwInterfaceSpanDumpClient struct {
	api.Stream
}

func (c *serviceClient_SwInterfaceSpanDumpClient) Recv() (*SwInterfaceSpanDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SwInterfaceSpanDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) SwInterfaceSpanEnableDisable(ctx context.Context, in *SwInterfaceSpanEnableDisable) (*SwInterfaceSpanEnableDisableReply, error) {
	out := new(SwInterfaceSpanEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

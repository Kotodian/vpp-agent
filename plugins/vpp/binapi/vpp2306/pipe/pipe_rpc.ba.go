// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package pipe

import (
	"context"
	"fmt"
	"io"

	api "go.fd.io/govpp/api"
	memclnt "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2306/memclnt"
)

// RPCService defines RPC service pipe.
type RPCService interface {
	PipeCreate(ctx context.Context, in *PipeCreate) (*PipeCreateReply, error)
	PipeDelete(ctx context.Context, in *PipeDelete) (*PipeDeleteReply, error)
	PipeDump(ctx context.Context, in *PipeDump) (RPCService_PipeDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) PipeCreate(ctx context.Context, in *PipeCreate) (*PipeCreateReply, error) {
	out := new(PipeCreateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PipeDelete(ctx context.Context, in *PipeDelete) (*PipeDeleteReply, error) {
	out := new(PipeDeleteReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PipeDump(ctx context.Context, in *PipeDump) (RPCService_PipeDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_PipeDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_PipeDumpClient interface {
	Recv() (*PipeDetails, error)
	api.Stream
}

type serviceClient_PipeDumpClient struct {
	api.Stream
}

func (c *serviceClient_PipeDumpClient) Recv() (*PipeDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *PipeDetails:
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

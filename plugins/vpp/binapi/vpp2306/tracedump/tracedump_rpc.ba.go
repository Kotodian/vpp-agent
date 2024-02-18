// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package tracedump

import (
	"context"
	"fmt"
	"io"

	api "go.fd.io/govpp/api"
	memclnt "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2306/memclnt"
)

// RPCService defines RPC service tracedump.
type RPCService interface {
	TraceCapturePackets(ctx context.Context, in *TraceCapturePackets) (*TraceCapturePacketsReply, error)
	TraceClearCache(ctx context.Context, in *TraceClearCache) (*TraceClearCacheReply, error)
	TraceClearCapture(ctx context.Context, in *TraceClearCapture) (*TraceClearCaptureReply, error)
	TraceDump(ctx context.Context, in *TraceDump) (RPCService_TraceDumpClient, error)
	TraceSetFilters(ctx context.Context, in *TraceSetFilters) (*TraceSetFiltersReply, error)
	TraceV2Dump(ctx context.Context, in *TraceV2Dump) (RPCService_TraceV2DumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) TraceCapturePackets(ctx context.Context, in *TraceCapturePackets) (*TraceCapturePacketsReply, error) {
	out := new(TraceCapturePacketsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) TraceClearCache(ctx context.Context, in *TraceClearCache) (*TraceClearCacheReply, error) {
	out := new(TraceClearCacheReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) TraceClearCapture(ctx context.Context, in *TraceClearCapture) (*TraceClearCaptureReply, error) {
	out := new(TraceClearCaptureReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) TraceDump(ctx context.Context, in *TraceDump) (RPCService_TraceDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_TraceDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_TraceDumpClient interface {
	Recv() (*TraceDetails, *TraceDumpReply, error)
	api.Stream
}

type serviceClient_TraceDumpClient struct {
	api.Stream
}

func (c *serviceClient_TraceDumpClient) Recv() (*TraceDetails, *TraceDumpReply, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, nil, err
	}
	switch m := msg.(type) {
	case *TraceDetails:
		return m, nil, nil
	case *TraceDumpReply:
		if err := api.RetvalToVPPApiError(m.Retval); err != nil {
			c.Stream.Close()
			return nil, m, err
		}
		err = c.Stream.Close()
		if err != nil {
			return nil, m, err
		}
		return nil, m, io.EOF
	default:
		return nil, nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) TraceSetFilters(ctx context.Context, in *TraceSetFilters) (*TraceSetFiltersReply, error) {
	out := new(TraceSetFiltersReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) TraceV2Dump(ctx context.Context, in *TraceV2Dump) (RPCService_TraceV2DumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_TraceV2DumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_TraceV2DumpClient interface {
	Recv() (*TraceV2Details, error)
	api.Stream
}

type serviceClient_TraceV2DumpClient struct {
	api.Stream
}

func (c *serviceClient_TraceV2DumpClient) Recv() (*TraceV2Details, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *TraceV2Details:
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

// Code generated by goa v3.15.2, DO NOT EDIT.
//
// messages gRPC client
//
// Command:
// $ goa gen goa-grpc-web/design

package client

import (
	"context"
	messagespb "goa-grpc-web/gen/grpc/messages/pb"
	messages "goa-grpc-web/gen/messages"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli messagespb.MessagesClient
	opts    []grpc.CallOption
}

// SubscribeClientStream implements the messages.SubscribeClientStream
// interface.
type SubscribeClientStream struct {
	stream messagespb.Messages_SubscribeClient
}

// NewClient instantiates gRPC client for all the messages service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: messagespb.NewMessagesClient(cc),
		opts:    opts,
	}
}

// Send calls the "Send" function in messagespb.MessagesClient interface.
func (c *Client) Send() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildSendFunc(c.grpccli, c.opts...),
			EncodeSendRequest,
			nil)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// Subscribe calls the "Subscribe" function in messagespb.MessagesClient
// interface.
func (c *Client) Subscribe() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildSubscribeFunc(c.grpccli, c.opts...),
			nil,
			DecodeSubscribeResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// Recv reads instances of "messagespb.SubscribeResponse" from the "subscribe"
// endpoint gRPC stream.
func (s *SubscribeClientStream) Recv() (*messages.Message, error) {
	var res *messages.Message
	v, err := s.stream.Recv()
	if err != nil {
		return res, err
	}
	return NewSubscribeResponseMessage_(v), nil
}

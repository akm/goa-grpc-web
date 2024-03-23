// Code generated with goa v3.15.2, DO NOT EDIT.
//
// messages protocol buffer definition
//
// Command:
// $ goa gen goa-grpc-web/design

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: goagen_goa-grpc-web_messages.proto

package messagespb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Messages_Send_FullMethodName      = "/messages.Messages/Send"
	Messages_Subscribe_FullMethodName = "/messages.Messages/Subscribe"
)

// MessagesClient is the client API for Messages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagesClient interface {
	// Send implements send.
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	// Subscribe to events sent such new chat messages.
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Messages_SubscribeClient, error)
}

type messagesClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagesClient(cc grpc.ClientConnInterface) MessagesClient {
	return &messagesClient{cc}
}

func (c *messagesClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, Messages_Send_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Messages_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Messages_ServiceDesc.Streams[0], Messages_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messagesSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Messages_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type messagesSubscribeClient struct {
	grpc.ClientStream
}

func (x *messagesSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessagesServer is the server API for Messages service.
// All implementations must embed UnimplementedMessagesServer
// for forward compatibility
type MessagesServer interface {
	// Send implements send.
	Send(context.Context, *SendRequest) (*SendResponse, error)
	// Subscribe to events sent such new chat messages.
	Subscribe(*SubscribeRequest, Messages_SubscribeServer) error
	mustEmbedUnimplementedMessagesServer()
}

// UnimplementedMessagesServer must be embedded to have forward compatible implementations.
type UnimplementedMessagesServer struct {
}

func (UnimplementedMessagesServer) Send(context.Context, *SendRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedMessagesServer) Subscribe(*SubscribeRequest, Messages_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedMessagesServer) mustEmbedUnimplementedMessagesServer() {}

// UnsafeMessagesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagesServer will
// result in compilation errors.
type UnsafeMessagesServer interface {
	mustEmbedUnimplementedMessagesServer()
}

func RegisterMessagesServer(s grpc.ServiceRegistrar, srv MessagesServer) {
	s.RegisterService(&Messages_ServiceDesc, srv)
}

func _Messages_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messages_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessagesServer).Subscribe(m, &messagesSubscribeServer{stream})
}

type Messages_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type messagesSubscribeServer struct {
	grpc.ServerStream
}

func (x *messagesSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Messages_ServiceDesc is the grpc.ServiceDesc for Messages service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messages_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messages.Messages",
	HandlerType: (*MessagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Messages_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Messages_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "goagen_goa-grpc-web_messages.proto",
}

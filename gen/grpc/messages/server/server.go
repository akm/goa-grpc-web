// Code generated by goa v3.15.2, DO NOT EDIT.
//
// messages gRPC server
//
// Command:
// $ goa gen goa-grpc-web/design

package server

import (
	"context"
	messagespb "goa-grpc-web/gen/grpc/messages/pb"
	messages "goa-grpc-web/gen/messages"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the messagespb.MessagesServer interface.
type Server struct {
	SendH      goagrpc.UnaryHandler
	SubscribeH goagrpc.StreamHandler
	messagespb.UnimplementedMessagesServer
}

// SubscribeServerStream implements the messages.SubscribeServerStream
// interface.
type SubscribeServerStream struct {
	stream messagespb.Messages_SubscribeServer
}

// New instantiates the server struct with the messages service endpoints.
func New(e *messages.Endpoints, uh goagrpc.UnaryHandler, sh goagrpc.StreamHandler) *Server {
	return &Server{
		SendH:      NewSendHandler(e.Send, uh),
		SubscribeH: NewSubscribeHandler(e.Subscribe, sh),
	}
}

// NewSendHandler creates a gRPC handler which serves the "messages" service
// "send" endpoint.
func NewSendHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeSendRequest, EncodeSendResponse)
	}
	return h
}

// Send implements the "Send" method in messagespb.MessagesServer interface.
func (s *Server) Send(ctx context.Context, message *messagespb.SendRequest) (*messagespb.SendResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "send")
	ctx = context.WithValue(ctx, goa.ServiceKey, "messages")
	resp, err := s.SendH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*messagespb.SendResponse), nil
}

// NewSubscribeHandler creates a gRPC handler which serves the "messages"
// service "subscribe" endpoint.
func NewSubscribeHandler(endpoint goa.Endpoint, h goagrpc.StreamHandler) goagrpc.StreamHandler {
	if h == nil {
		h = goagrpc.NewStreamHandler(endpoint, nil)
	}
	return h
}

// Subscribe implements the "Subscribe" method in messagespb.MessagesServer
// interface.
func (s *Server) Subscribe(message *messagespb.SubscribeRequest, stream messagespb.Messages_SubscribeServer) error {
	ctx := stream.Context()
	ctx = context.WithValue(ctx, goa.MethodKey, "subscribe")
	ctx = context.WithValue(ctx, goa.ServiceKey, "messages")
	_, err := s.SubscribeH.Decode(ctx, message)
	if err != nil {
		return goagrpc.EncodeError(err)
	}
	ep := &messages.SubscribeEndpointInput{
		Stream: &SubscribeServerStream{stream: stream},
	}
	err = s.SubscribeH.Handle(ctx, ep)
	if err != nil {
		return goagrpc.EncodeError(err)
	}
	return nil
}

// Send streams instances of "messagespb.SubscribeResponse" to the "subscribe"
// endpoint gRPC stream.
func (s *SubscribeServerStream) Send(res *messages.Message) error {
	v := NewProtoMessage_SubscribeResponse(res)
	return s.stream.Send(v)
}

func (s *SubscribeServerStream) Close() error {
	// nothing to do here
	return nil
}
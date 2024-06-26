// Code generated by goa v3.15.2, DO NOT EDIT.
//
// messages gRPC server encoders and decoders
//
// Command:
// $ goa gen goa-grpc-web/design

package server

import (
	"context"
	messagespb "goa-grpc-web/gen/grpc/messages/pb"
	messages "goa-grpc-web/gen/messages"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeSendResponse encodes responses from the "messages" service "send"
// endpoint.
func EncodeSendResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	resp := NewProtoSendResponse()
	return resp, nil
}

// DecodeSendRequest decodes requests sent to "messages" service "send"
// endpoint.
func DecodeSendRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *messagespb.SendRequest
		ok      bool
	)
	{
		if message, ok = v.(*messagespb.SendRequest); !ok {
			return nil, goagrpc.ErrInvalidType("messages", "send", "*messagespb.SendRequest", v)
		}
	}
	var payload *messages.Message
	{
		payload = NewSendPayload(message)
	}
	return payload, nil
}

// EncodeSubscribeResponse encodes responses from the "messages" service
// "subscribe" endpoint.
func EncodeSubscribeResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	result, ok := v.(*messages.Message)
	if !ok {
		return nil, goagrpc.ErrInvalidType("messages", "subscribe", "*messages.Message", v)
	}
	resp := NewProtoSubscribeResponse(result)
	return resp, nil
}

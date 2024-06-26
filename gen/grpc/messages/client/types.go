// Code generated by goa v3.15.2, DO NOT EDIT.
//
// messages gRPC client types
//
// Command:
// $ goa gen goa-grpc-web/design

package client

import (
	messagespb "goa-grpc-web/gen/grpc/messages/pb"
	messages "goa-grpc-web/gen/messages"
)

// NewProtoSendRequest builds the gRPC request type from the payload of the
// "send" endpoint of the "messages" service.
func NewProtoSendRequest(payload *messages.Message) *messagespb.SendRequest {
	message := &messagespb.SendRequest{
		Message_: payload.Message,
	}
	return message
}

// NewProtoSubscribeRequest builds the gRPC request type from the payload of
// the "subscribe" endpoint of the "messages" service.
func NewProtoSubscribeRequest() *messagespb.SubscribeRequest {
	message := &messagespb.SubscribeRequest{}
	return message
}

func NewSubscribeResponseMessage_(v *messagespb.SubscribeResponse) *messages.Message {
	result := &messages.Message{
		Message: v.Message_,
	}
	return result
}

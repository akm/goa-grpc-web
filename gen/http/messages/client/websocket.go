// Code generated by goa v3.15.2, DO NOT EDIT.
//
// messages WebSocket client streaming
//
// Command:
// $ goa gen goa-grpc-web/design

package client

import (
	messages "goa-grpc-web/gen/messages"
	"io"

	"github.com/gorilla/websocket"
	goahttp "goa.design/goa/v3/http"
)

// ConnConfigurer holds the websocket connection configurer functions for the
// streaming endpoints in "messages" service.
type ConnConfigurer struct {
	SubscribeFn goahttp.ConnConfigureFunc
}

// SubscribeClientStream implements the messages.SubscribeClientStream
// interface.
type SubscribeClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// NewConnConfigurer initializes the websocket connection configurer function
// with fn for all the streaming endpoints in "messages" service.
func NewConnConfigurer(fn goahttp.ConnConfigureFunc) *ConnConfigurer {
	return &ConnConfigurer{
		SubscribeFn: fn,
	}
}

// Recv reads instances of "messages.Message" from the "subscribe" endpoint
// websocket connection.
func (s *SubscribeClientStream) Recv() (*messages.Message, error) {
	var (
		rv   *messages.Message
		body SubscribeResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	err = ValidateSubscribeResponseBody(&body)
	if err != nil {
		return rv, err
	}
	res := NewSubscribeMessageOK(&body)
	return res, nil
}
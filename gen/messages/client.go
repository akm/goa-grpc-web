// Code generated by goa v3.15.2, DO NOT EDIT.
//
// messages client
//
// Command:
// $ goa gen goa-grpc-web/design

package messages

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "messages" service client.
type Client struct {
	SendEndpoint      goa.Endpoint
	SubscribeEndpoint goa.Endpoint
}

// NewClient initializes a "messages" service client given the endpoints.
func NewClient(send, subscribe goa.Endpoint) *Client {
	return &Client{
		SendEndpoint:      send,
		SubscribeEndpoint: subscribe,
	}
}

// Send calls the "send" endpoint of the "messages" service.
func (c *Client) Send(ctx context.Context, p *Message) (err error) {
	_, err = c.SendEndpoint(ctx, p)
	return
}

// Subscribe calls the "subscribe" endpoint of the "messages" service.
func (c *Client) Subscribe(ctx context.Context) (res SubscribeClientStream, err error) {
	var ires any
	ires, err = c.SubscribeEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(SubscribeClientStream), nil
}

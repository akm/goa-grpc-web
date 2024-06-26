// Code generated by goa v3.15.2, DO NOT EDIT.
//
// messages client HTTP transport
//
// Command:
// $ goa gen goa-grpc-web/design

package client

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the messages service endpoint HTTP clients.
type Client struct {
	// Send Doer is the HTTP client used to make requests to the send endpoint.
	SendDoer goahttp.Doer

	// Subscribe Doer is the HTTP client used to make requests to the subscribe
	// endpoint.
	SubscribeDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme     string
	host       string
	encoder    func(*http.Request) goahttp.Encoder
	decoder    func(*http.Response) goahttp.Decoder
	dialer     goahttp.Dialer
	configurer *ConnConfigurer
}

// NewClient instantiates HTTP clients for all the messages service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
	dialer goahttp.Dialer,
	cfn *ConnConfigurer,
) *Client {
	if cfn == nil {
		cfn = &ConnConfigurer{}
	}
	return &Client{
		SendDoer:            doer,
		SubscribeDoer:       doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
		dialer:              dialer,
		configurer:          cfn,
	}
}

// Send returns an endpoint that makes HTTP requests to the messages service
// send server.
func (c *Client) Send() goa.Endpoint {
	var (
		encodeRequest  = EncodeSendRequest(c.encoder)
		decodeResponse = DecodeSendResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildSendRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SendDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("messages", "send", err)
		}
		return decodeResponse(resp)
	}
}

// Subscribe returns an endpoint that makes HTTP requests to the messages
// service subscribe server.
func (c *Client) Subscribe() goa.Endpoint {
	var (
		decodeResponse = DecodeSubscribeResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildSubscribeRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("messages", "subscribe", err)
		}
		if c.configurer.SubscribeFn != nil {
			var cancel context.CancelFunc
			ctx, cancel = context.WithCancel(ctx)
			conn = c.configurer.SubscribeFn(conn, cancel)
		}
		go func() {
			<-ctx.Done()
			conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, "client closing connection"),
				time.Now().Add(time.Second),
			)
			conn.Close()
		}()
		stream := &SubscribeClientStream{conn: conn}
		return stream, nil
	}
}

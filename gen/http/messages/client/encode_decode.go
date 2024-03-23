// Code generated by goa v3.15.2, DO NOT EDIT.
//
// messages HTTP client encoders and decoders
//
// Command:
// $ goa gen goa-grpc-web/design

package client

import (
	"bytes"
	"context"
	messages "goa-grpc-web/gen/messages"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildSendRequest instantiates a HTTP request object with method and path set
// to call the "messages" service "send" endpoint
func (c *Client) BuildSendRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SendMessagesPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("messages", "send", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSendRequest returns an encoder for requests sent to the messages send
// server.
func EncodeSendRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*messages.Message)
		if !ok {
			return goahttp.ErrInvalidType("messages", "send", "*messages.Message", v)
		}
		body := NewSendRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("messages", "send", err)
		}
		return nil
	}
}

// DecodeSendResponse returns a decoder for responses returned by the messages
// send endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeSendResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("messages", "send", resp.StatusCode, string(body))
		}
	}
}

// BuildSubscribeRequest instantiates a HTTP request object with method and
// path set to call the "messages" service "subscribe" endpoint
func (c *Client) BuildSubscribeRequest(ctx context.Context, v any) (*http.Request, error) {
	scheme := c.scheme
	switch c.scheme {
	case "http":
		scheme = "ws"
	case "https":
		scheme = "wss"
	}
	u := &url.URL{Scheme: scheme, Host: c.host, Path: SubscribeMessagesPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("messages", "subscribe", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeSubscribeResponse returns a decoder for responses returned by the
// messages subscribe endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeSubscribeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body SubscribeResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("messages", "subscribe", err)
			}
			err = ValidateSubscribeResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("messages", "subscribe", err)
			}
			res := NewSubscribeMessageOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("messages", "subscribe", resp.StatusCode, string(body))
		}
	}
}

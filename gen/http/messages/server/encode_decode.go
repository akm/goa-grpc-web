// Code generated by goa v3.15.2, DO NOT EDIT.
//
// messages HTTP server encoders and decoders
//
// Command:
// $ goa gen goa-grpc-web/design

package server

import (
	"context"
	"io"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeSendResponse returns an encoder for responses returned by the messages
// send endpoint.
func EncodeSendResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusCreated)
		return nil
	}
}

// DecodeSendRequest returns a decoder for requests sent to the messages send
// endpoint.
func DecodeSendRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body SendRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateSendRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewSendMessage(&body)

		return payload, nil
	}
}
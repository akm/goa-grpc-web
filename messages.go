package notification

import (
	"context"
	messages "goa-grpc-web/gen/messages"
	"log"
)

// messages service example implementation.
// The example methods log the requests and return zero values.
type messagessrvc struct {
	logger *log.Logger
}

// NewMessages returns the messages service implementation.
func NewMessages(logger *log.Logger) messages.Service {
	return &messagessrvc{logger}
}

// Send implements send.
func (s *messagessrvc) Send(ctx context.Context, p *messages.Message) (err error) {
	s.logger.Print("messages.send")
	return
}

// Subscribe to events sent such new chat messages.
func (s *messagessrvc) Subscribe(ctx context.Context, stream messages.SubscribeServerStream) (err error) {
	s.logger.Print("messages.subscribe")
	return
}

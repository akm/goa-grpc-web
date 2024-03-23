package notification

import (
	"context"
	messages "goa-grpc-web/gen/messages"
	"log"
	"time"
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

var messageChannels = []chan messages.Message{}

// Send implements send.
func (s *messagessrvc) Send(ctx context.Context, p *messages.Message) (err error) {
	s.logger.Print("messages.send")

	go func() {
		for _, ch := range messageChannels {
			ch <- *p
		}
	}()

	return
}

// Subscribe to events sent such new chat messages.
func (s *messagessrvc) Subscribe(ctx context.Context, stream messages.SubscribeServerStream) (err error) {
	s.logger.Printf("messages.subscribe\n")

	defer func() {
		if err := stream.Close(); err != nil {
			s.logger.Printf("failed to close stream: [%T] %+v", err, err)
		} else {
			s.logger.Printf("Stream Closed Successfully")
		}
	}()

	ch := make(chan messages.Message)
	messageChannels = append(messageChannels, ch)

	defer func() {
		for i, c := range messageChannels {
			if c == ch {
				messageChannels = append(messageChannels[:i], messageChannels[i+1:]...)
				close(ch)
				break
			}
		}
	}()

	interval := 500 * time.Millisecond
	ticker := time.NewTicker(interval)

	done := false
	for {
		select {
		case <-ctx.Done():
			done = true
		case <-ticker.C:
			msg := <-ch

			if err := stream.Send(&msg); err != nil {
				return err
			}
		}
		if done {
			break
		}
	}

	return
}

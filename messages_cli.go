package notification

import (
	"fmt"
	"io"

	messagesvc "goa-grpc-web/gen/messages"
)

// https://github.com/goadesign/examples/blob/master/streaming/chattercli.go
func InteractWithStreams(data interface{}) {
	if data == nil {
		return
	}

	switch stream := data.(type) {
	case messagesvc.SubscribeClientStream:
		for {
			p, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(fmt.Errorf("error reading from stream: %v", err))
				break
			}
			fmt.Printf("Received message: %+v\n", *p)
		}
	}
}

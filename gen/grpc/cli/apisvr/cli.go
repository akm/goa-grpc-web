// Code generated by goa v3.15.2, DO NOT EDIT.
//
// apisvr gRPC client CLI support package
//
// Command:
// $ goa gen goa-grpc-web/design

package cli

import (
	"flag"
	"fmt"
	messagesc "goa-grpc-web/gen/grpc/messages/client"
	"os"

	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `messages (send|subscribe)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` messages send --message '{
      "message": "Sint id soluta quos animi architecto libero."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, any, error) {
	var (
		messagesFlags = flag.NewFlagSet("messages", flag.ContinueOnError)

		messagesSendFlags       = flag.NewFlagSet("send", flag.ExitOnError)
		messagesSendMessageFlag = messagesSendFlags.String("message", "", "")

		messagesSubscribeFlags = flag.NewFlagSet("subscribe", flag.ExitOnError)
	)
	messagesFlags.Usage = messagesUsage
	messagesSendFlags.Usage = messagesSendUsage
	messagesSubscribeFlags.Usage = messagesSubscribeUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "messages":
			svcf = messagesFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "messages":
			switch epn {
			case "send":
				epf = messagesSendFlags

			case "subscribe":
				epf = messagesSubscribeFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "messages":
			c := messagesc.NewClient(cc, opts...)
			switch epn {
			case "send":
				endpoint = c.Send()
				data, err = messagesc.BuildSendPayload(*messagesSendMessageFlag)
			case "subscribe":
				endpoint = c.Subscribe()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// messagesUsage displays the usage of the messages command and its subcommands.
func messagesUsage() {
	fmt.Fprintf(os.Stderr, `Service is the messages service interface.
Usage:
    %[1]s [globalflags] messages COMMAND [flags]

COMMAND:
    send: Send implements send.
    subscribe: Subscribe to events sent such new chat messages.

Additional help:
    %[1]s messages COMMAND --help
`, os.Args[0])
}
func messagesSendUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] messages send -message JSON

Send implements send.
    -message JSON: 

Example:
    %[1]s messages send --message '{
      "message": "Sint id soluta quos animi architecto libero."
   }'
`, os.Args[0])
}

func messagesSubscribeUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] messages subscribe

Subscribe to events sent such new chat messages.

Example:
    %[1]s messages subscribe
`, os.Args[0])
}

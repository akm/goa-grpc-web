package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("messages", func() {
	HTTP(func() {
		Path("/messages")
	})

	GRPC(func() {
	})

	Method("send", func() {
		Payload(MessageType)
		HTTP(func() {
			POST("")
			Response(StatusCreated)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("subscribe", func() {
		Description("Subscribe to events sent such new chat messages.")

		StreamingResult(MessageType)

		HTTP(func() {
			GET("/subscribe")
			// httpIdToken()
			Response(StatusOK)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})
})

var MessageType = Type("Message", func() {
	Field(1, "message", String)
	Required("message")
})

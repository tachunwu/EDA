package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))
	select{}
}
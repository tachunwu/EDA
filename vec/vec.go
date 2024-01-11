package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)
var servers =`nats://localhost:5001
			  `

func main() {
	nc, _ := nats.Connect(servers)
	nc.Subscribe("vec", func(m *nats.Msg){
		nc.Publish("vec.out", []byte("Data"))
	})
	nc.Subscribe("vec", func(m *nats.Msg){
		nc.Publish("vec.out", []byte("Data"))
	})
	nc.Subscribe("vec", func(m *nats.Msg){
		nc.Publish("vec.out", []byte("Data"))
	})

	// Client start!

	// Time based vec io
	sub, _ := nc.SubscribeSync("vec.out")
	nc.Flush()

	nc.Publish("vec",[]byte("vec_io"))

	max := 100 * time.Millisecond
	minResponses := 2
	start := time.Now()
	responses := []string{}
	for time.Now().Sub(start) < max {
		msg, err := sub.NextMsg(1 * time.Second)
		if err != nil {
			break
		}
	
		responses = append(responses, string(msg.Data))

		// By length
		if len(responses) >= minResponses {
			break
		}
	}
	sub.Unsubscribe()
	fmt.Println(len(responses))

	select {}
}
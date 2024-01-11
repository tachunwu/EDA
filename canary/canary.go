package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nats-io/nats.go"
)
var servers =`nats://localhost:5001
			  `

func main() {
	nc, _ := nats.Connect(servers)
	nc.QueueSubscribe("canary", "CanaryPool", func(m *nats.Msg){
		if rand.Intn(10) < 6 {
			panic("BANG!")
		}
		nc.Publish("canary.out", []byte("Data"))
	})

	// Client start!

	// Time based vec io
	sub, _ := nc.SubscribeSync("canary.out")
	nc.Flush()

	nc.Publish("canary",[]byte("CCCanary!"))

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
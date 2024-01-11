package main

import (
	"math/rand"

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

		m.Respond([]byte("Data"))
	})

	select {}
}
package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

var servers =`nats://localhost:5001
			  `

// nats publish FT.pool "Yoyoyo" -s localhost:5001
// nats publish FT.pool "Yoyoyo" -s localhost:5001
// kill
// nats publish FT.pool "Yoyoyo" -s localhost:5001
// nats publish FT.pool "Yoyoyo" -s localhost:5001

func main() {

	nc, _ := nats.Connect(servers)
	nc.QueueSubscribe("FT.pool","Pool", func(m *nats.Msg){
		fmt.Println("Fault-Tolerance Group:", string(m.Data))
	})
	select{}
}
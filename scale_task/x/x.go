package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

var servers =`nats://localhost:5000, nats://localhost:5001, nats://localhost:5002,
		      nats://localhost:5003, nats://localhost:5004, nats://localhost:5005,
			  nats://localhost:5006, nats://localhost:5007, nats://localhost:5008,nats://localhost:5009
			  `

// nats publish task.pool "Yoyoyo" -s localhost:5000,localhost:5001,localhost:5002,localhost:5003,localhost:5004,localhost:5005,localhost:5006,localhost:5007,localhost:5008,localhost:5009
func main() {

	nc, _ := nats.Connect(servers)
	nc.QueueSubscribe("task.pool","pool",func(m *nats.Msg){
		fmt.Println("Process Task:", string(m.Data))
	})
	nc.QueueSubscribe("task.pool","pool",func(m *nats.Msg){
		fmt.Println("Process Task:", string(m.Data))
	})
	nc.QueueSubscribe("task.pool","pool",func(m *nats.Msg){
		fmt.Println("Process Task:", string(m.Data))
	})
	nc.QueueSubscribe("task.pool","pool",func(m *nats.Msg){
		fmt.Println("Process Task:", string(m.Data))
	})
	nc.QueueSubscribe("task.pool","pool",func(m *nats.Msg){
		fmt.Println("Process Task:", string(m.Data))
	})
	nc.QueueSubscribe("task.pool","pool",func(m *nats.Msg){
		fmt.Println("Process Task:", string(m.Data))
	})
	select{}
}
package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

var servers =`nats://localhost:5000, nats://localhost:5001, nats://localhost:5002,
		      nats://localhost:5003, nats://localhost:5004, nats://localhost:5005,
			  nats://localhost:5006, nats://localhost:5007, nats://localhost:5008,nats://localhost:5009
			  `

// nats publish function "Yoyoyo" -s localhost:5000,localhost:5001,localhost:5002,localhost:5003,localhost:5004,localhost:5005,localhost:5006,localhost:5007,localhost:5008,localhost:5009
func main() {

	nc, _ := nats.Connect(servers)
	nc.QueueSubscribe("function.A","A",func(m *nats.Msg){
		fmt.Println("A Process Task:", string(m.Data))
	})
	nc.QueueSubscribe("function.B","B",func(m *nats.Msg){
		fmt.Println("B Process Task:", string(m.Data))
	})
	nc.QueueSubscribe("function.C","C",func(m *nats.Msg){
		fmt.Println("C Process Task:", string(m.Data))
	})
	nc.QueueSubscribe("function.D","D",func(m *nats.Msg){
		fmt.Println("D Process Task:", string(m.Data))
	})
	nc.QueueSubscribe("function.E","E",func(m *nats.Msg){
		fmt.Println("E Process Task:", string(m.Data))
	})
	nc.QueueSubscribe("function.F","F",func(m *nats.Msg){
		fmt.Println("F Process Task:", string(m.Data))
	})
	select{}
}
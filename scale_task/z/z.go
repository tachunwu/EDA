package main

import (
	"fmt"
	"hash/fnv"
	"strconv"

	"github.com/nats-io/nats.go"
)

var servers =`nats://localhost:5000, nats://localhost:5001, nats://localhost:5002,
		      nats://localhost:5003, nats://localhost:5004, nats://localhost:5005,
			  nats://localhost:5006, nats://localhost:5007, nats://localhost:5008,nats://localhost:5009
			  `

// nats publish function "Yoyoyo" -s localhost:5000,localhost:5001,localhost:5002,localhost:5003,localhost:5004,localhost:5005,localhost:5006,localhost:5007,localhost:5008,localhost:5009
func main() {

	nc, _ := nats.Connect(servers)
	nc.QueueSubscribe("Partition.0","A",func(m *nats.Msg){
		fmt.Println("Partition-0 Store Data:", string(m.Data))
	})
	nc.QueueSubscribe("Partition.1","B",func(m *nats.Msg){
		fmt.Println("Partition-1 Store Data:", string(m.Data))
	})
	nc.QueueSubscribe("Partition.2","C",func(m *nats.Msg){
		fmt.Println("Partition-2 Store Data:", string(m.Data))
	})
	nc.QueueSubscribe("Partition.3","D",func(m *nats.Msg){
		fmt.Println("Partition-3 Store Data:", string(m.Data))
	})
	nc.QueueSubscribe("Partition.4","E",func(m *nats.Msg){
		fmt.Println("Partition-4 Store Data:", string(m.Data))
	})
	nc.QueueSubscribe("Partition.5","F",func(m *nats.Msg){
		fmt.Println("Partition-5 Store Data:", string(m.Data))
	})
	i:=0
	// i: Index for the data
	// strconv(i)
	for i < 100 {
		partitionId :=hashStringToInt(strconv.Itoa(i)) % 6
		nc.Publish("Partition." + strconv.Itoa(int(partitionId)),[]byte("Data"))
		i++
	}
	
	select{}
}

func hashStringToInt(s string) uint32 {
    h := fnv.New32a()
    h.Write([]byte(s))
    return h.Sum32()
}
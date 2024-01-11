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

	maxFailCount := 3
	failCount := 0
	for failCount < maxFailCount{
		msg, err := nc.Request("canary", []byte("C"), 1 * time.Second)
		if err != nil {
			failCount++
			fmt.Println("Fail:",failCount)
			continue
		}
		fmt.Println("Canary Get! At round:",failCount+1,"Response data:", string(msg.Data))
		break
	}


}
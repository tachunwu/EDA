package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	eda "github.com/tachunwu/EDA/protos"
	"google.golang.org/protobuf/proto"
)

func main() {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)
	task := &eda.Task{
		Delay: 10,
	}
	taskData, _ := proto.Marshal(task)

	nc.Subscribe("task.sync", TaskProcessor)
	// nc.Publish("task.sync", taskData)
	rep, _:= nc.Request("task.sync", taskData, 1 *time.Second)
	doned := &eda.TaskDoned{}
	proto.Unmarshal(rep.Data, doned)
	fmt.Println("Task Doned!", doned)
	select{}

}


func TaskProcessor(m *nats.Msg) {
	newTask := &eda.Task{}
	proto.Unmarshal(m.Data, newTask)
	time.Sleep(time.Millisecond * time.Duration(newTask.Delay))
	fmt.Println("Task process: ", newTask.Delay," ms")
	
	doned := &eda.TaskDoned{}
	b,_:= proto.Marshal(doned)
	m.Respond(b)
}
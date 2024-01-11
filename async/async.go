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
	defer nc.Close()

	// Prepare task
	task := &eda.Task{
		Delay: 1000 * 10,
	}
	taskData, _ := proto.Marshal(task)

	// Start processor
	nc.Subscribe("task.async", TaskProcessor) // task.async. + 'task UUIDv4'

	// Submit task
	// nc.Publish("task.async", taskData)

	// Check after Enqueued
	status := &eda.TaskStatus{}

	rep, _:= nc.Request("task.async", taskData, 1 *time.Second)
	proto.Unmarshal(rep.Data, status)
	fmt.Println("[Application] Task status:", status.State)

	// Wait processor
	time.Sleep(15* time.Second)

	// Check after 15 seconds
	rep, _= nc.Request("task.async", taskData, 1 *time.Second)
	proto.Unmarshal(rep.Data, status)
	fmt.Println("[Application] Task status:", status.State)
	select{}

}

var IsDoned = false

func TaskProcessor(m *nats.Msg) {
	fmt.Println("[Processor]: Is Task done?", IsDoned)
	newTask := &eda.Task{}
	proto.Unmarshal(m.Data, newTask)

	if !IsDoned {
		go func(){
			time.Sleep(time.Millisecond * time.Duration(newTask.Delay))
			fmt.Println("[Processor]: ", newTask.Delay," ms")
			IsDoned = true
			doned := &eda.TaskStatus{
				State: eda.TaskState_DONED,
			}
			b,_:= proto.Marshal(doned)
			m.Respond(b)
		}()
	
		queued := &eda.TaskStatus{
			State: eda.TaskState_QUEUED,
		}
		b,_:= proto.Marshal(queued)
		m.Respond(b)
	} else {
		doned := &eda.TaskStatus{
			State: eda.TaskState_DONED,
		}
		b,_:= proto.Marshal(doned)
		m.Respond(b)
	}

}
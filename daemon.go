package blacktable

import (
	"fmt"
	"reflect"
	"time"
)

type Task interface {
	Run(fromIp string, data []byte)
}

type PeriodicTask struct {
	Name   string
	Start  time.Time
	Period time.Duration
	Times  int // -1 = infinite
}

func (bt *BlackTable) startDaemon() {
	// fmt.Println("startDaemon")
	for {
		select {
		case task := <-bt.taskChan:
			bt.addTask(task)
		}
	}
}

func (bt *BlackTable) addTask(task interface{}) {
	fmt.Println("addTask", task, reflect.TypeOf(task).Name())
	switch task.(type) {
	case *SendTask:
		bt.addSendTask(task.(*SendTask))
	case *ReceiveTask:
		bt.addListenTask(task.(*ReceiveTask))
	}
}

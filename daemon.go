package blacktable

import (
	"fmt"
	"reflect"
	"time"
)

type Task interface {
	Run(fromIp string, data []byte)
	GetSchedule() *PeriodicTask
}

type PeriodicTask struct {
	Name   string
	Start  time.Time
	Period time.Duration
	Times  int // -1 = infinite
}

func (pt *PeriodicTask) GetSchedule() *PeriodicTask {
	return pt
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

func (bt *BlackTable) enqueueTask(t Task) {
	schedule := t.GetSchedule()
	if len(schedule.Name) > 0 {
		bt.Tasks[schedule.Name] = t
	}

}

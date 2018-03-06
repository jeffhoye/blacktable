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
	Name          string
	Start         time.Time
	Period        time.Duration
	Times         int // -1 = infinite
	NextExecution time.Time
}

func (pt *PeriodicTask) GetSchedule() *PeriodicTask {
	return pt
}

func (bt *BlackTable) startDaemon() {
	// fmt.Println("startDaemon")
	for {
		nextWakeDuration := time.Now().Sub(bt.NextWake)
		select {
		case task := <-bt.taskChan:
			bt.addTask(task)
		case <-time.After(nextWakeDuration):
			bt.runReadyTasks()
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
	s := t.GetSchedule()
	if len(s.Name) > 0 {
		// TODO: consider removing/dequeuing the old one
		bt.Tasks[s.Name] = t
	}

	if s.Times == 0 {
		return
	}

	s.NextExecution = s.Start
	if s.Period > 0 {
		for s.NextExecution.Before(time.Now()) {
			s.NextExecution = s.NextExecution.Add(s.Period)
		}
	}

	bt.NextWake = time.Now().Add(time.Minute)
	for _, t := range bt.Tasks {
		s := t.GetSchedule()
		if s.Times != 0 && s.NextExecution.Before(bt.NextWake) {
			bt.NextWake = s.NextExecution
		}
	}
}

func (bt *BlackTable) runReadyTasks() {
	now := time.Now()
	// requeue := make([]Task,0,0)
	runMe := make([]Task, 0, 0)
	for _, t := range bt.Tasks {
		s := t.GetSchedule()
		if s.Times != 0 {
			if !s.NextExecution.After(now) {
				s.Times--
				if s.Period > 0 {
					s.NextExecution = s.NextExecution.Add(s.Period)
				}
				runMe = append(runMe, t)
			}
		}
	}
	for _, t := range runMe {
		bt.enqueueTask(t)
	}
	for _, t := range runMe {
		go t.Run("", []byte{})
	}
}

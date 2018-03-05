package blacktable

import (
	"fmt"
	"time"
)

type PeriodicTask struct {
	Name   string
	Start  time.Time
	Period time.Duration
	Times  int // -1 = infinite
}

type NetworkMessage struct {
	PeriodicTask
	Protocol string // udp or tcp
	IpPort   string // ip address and port
	Message  []byte
}

func (nm *NetworkMessage) run(data []byte) {
	fullMessage := append(nm.Message, data...)
	fmt.Println(string(fullMessage))
}

func (bt *BlackTable) startDaemon() {
	select {
	case task := <-bt.taskChan:
		bt.addTask(task)
	}
}

func (bt *BlackTable) addTask(task interface{}) {
	switch task.(type) {
	case NetworkMessage:
		bt.addNetworkMessage(task.(*NetworkMessage))
	}
}

func (bt *BlackTable) addNetworkMessage(nm *NetworkMessage) {
	fmt.Println("Add Network Message")
}

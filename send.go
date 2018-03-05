package blacktable

import "fmt"

type SendTask struct {
	PeriodicTask
	Protocol string // udp or tcp
	IpPort   string // ip address and port
	Message  []byte
}

func (st *SendTask) run(fromIp string, data []byte) {
	fullMessage := append(st.Message, data...)
	fmt.Println(string(fullMessage))
}

func (bt *BlackTable) addSendTask(nm *SendTask) {
	fmt.Println("Add Send Task")
}

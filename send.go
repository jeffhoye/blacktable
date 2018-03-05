package blacktable

import (
	"fmt"
	"log"
	"net"
)

type SendTask struct {
	PeriodicTask
	Protocol string // udp or tcp
	ToIpPort string // ip address and port
	Message  []byte
}

func (st *SendTask) Run(fromIp string, data []byte) {
	fullMessage := append(st.Message, data...)
	fmt.Println("SendTask:",string(fullMessage))
	//Connect udp
	conn, err := net.Dial(st.Protocol, st.ToIpPort)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//simple Read
	// buffer := make([]byte, 1024)
	// conn.Read(buffer)

	//simple write
	conn.Write(fullMessage)
}

func (bt *BlackTable) addSendTask(st *SendTask) {
	fmt.Println("Add Send Task")
	// bt.Tasks[st.Name] = st

	go st.Run("", []byte(""))
}

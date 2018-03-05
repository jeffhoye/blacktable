package blacktable

import (
	"fmt"
	"log"
	"net"
)

type IpListener struct {
	Protocol string
	OnIpPort string
	Tasks    []*ListenTask
}

func (ipl *IpListener) addTask(lt *ListenTask) {
	ipl.Tasks = append(ipl.Tasks, lt)
}

func (bt *BlackTable) addIpListener(protocol string, onIpPort string) (ipl *IpListener) {
	ipl = &IpListener{
		Protocol: protocol,
		OnIpPort: onIpPort,
	}
	bt.IpListeners[protocol][onIpPort] = ipl
	go ipl.Start()
	return
}

func (ipl *IpListener) Start() {
	var err error
	switch ipl.Protocol {
	case "udp":
		err = ipl.startUDPserver()
	}
	if err != nil {
		log.Fatal("Error starting", ipl.Protocol, " listener on ", ipl.OnIpPort, err)
	}
}

func (ipl *IpListener) startUDPserver() error {
	udpAddr, err := net.ResolveUDPAddr(ipl.Protocol, ipl.OnIpPort)
	if err != nil {
		return err
	}
	udpConn, err := net.ListenUDP(ipl.Protocol, udpAddr)
	if err != nil {
		fmt.Println(err)
	}

	for {
		ipl.messageReceived(udpConn)
	}
	return nil
}

func (ipl *IpListener) messageReceived(conn *net.UDPConn) {
	var buf [2048]byte
	n, err := conn.Read(buf[0:])
	if err != nil {
		fmt.Println("Error Reading")
		return
	} else {
		fmt.Println("ListenTask", string(buf[0:n]))
		// fmt.Println("Package Done")
	}
}

func (bt *BlackTable) addListenTask(lt *ListenTask) {
	listener, ok := bt.IpListeners[lt.Protocol][lt.OnIpPort]
	if !ok {
		listener = bt.addIpListener(lt.Protocol, lt.OnIpPort)
	}
	listener.addTask(lt)
}

type ListenTask struct {
	PeriodicTask
	Protocol   string // udp or tcp
	OnIpPort   string
	FromIpPort string // regexp
	Message    string // regexp
}

func (lt *ListenTask) run(ip string, data []byte) {
	fmt.Println(string(data))
}

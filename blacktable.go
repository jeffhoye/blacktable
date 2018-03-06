package blacktable

import (
	"fmt"
	"sync"
	"time"
)

const HELP_TXT = `
How to use Black Table...
exit: to quit
`

type BlackTable struct {
	wg          *sync.WaitGroup
	taskChan    chan interface{}                  // new tasks get added here
	IpListeners map[string]map[string]*IpListener // protocol -> ipPort -> listener
	Tasks       map[string]Task                   // name -> task
	NextWake    time.Time
}

func NewBlackTable() (*BlackTable, error) {
	ipListeners := make(map[string]map[string]*IpListener, 0)
	ipListeners["udp"] = make(map[string]*IpListener, 0)
	ipListeners["tcp"] = make(map[string]*IpListener, 0)
	bt := &BlackTable{
		wg:          &sync.WaitGroup{},
		taskChan:    make(chan interface{}, 100),
		Tasks:       make(map[string]Task, 0),
		IpListeners: ipListeners,
	}
	return bt, nil
}

func (bt *BlackTable) Start() {
	fmt.Println("Running Blacktable")
	go bt.startDaemon()
	bt.wg.Add(1)
	go bt.readStdIn()
	bt.Wait()
}

func (bt *BlackTable) Wait() {
	bt.wg.Wait()
}

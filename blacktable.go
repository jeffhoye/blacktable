package blacktable

import (
	"fmt"
	"sync"
)

const HELP_TXT = `
How to use Black Table...
exit: to quit
`

type BlackTable struct {
	wg       *sync.WaitGroup
	taskChan chan interface{}
}

func NewBlackTable() (*BlackTable, error) {
	bt := &BlackTable{
		wg:       &sync.WaitGroup{},
		taskChan: make(chan interface{}, 100),
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

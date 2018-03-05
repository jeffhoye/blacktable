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
	wg *sync.WaitGroup
}

func NewBlackTable() (*BlackTable, error) {
	bt := &BlackTable{
		wg: &sync.WaitGroup{},
	}
	return bt, nil
}

func (bt *BlackTable) Start() {
	fmt.Println("Running Blacktable3")
	// go bt.runDaemon()
	bt.wg.Add(1)
	go bt.readStdIn()
	bt.Wait()
}

func (bt *BlackTable) Wait() {
	bt.wg.Wait()
}

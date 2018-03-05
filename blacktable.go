package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
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

func (bt *BlackTable) Run() {
	fmt.Println("Running Blacktable")
	// go bt.runDaemon()
	bt.wg.Add(1)
	go bt.readStdIn()
	bt.Wait()
}

func (bt *BlackTable) Wait() {
	bt.wg.Wait()
}

func newCsvReader(r io.Reader) *csv.Reader {
	cr := csv.NewReader(r)
	cr.FieldsPerRecord = -1
	return cr
}

func (bt *BlackTable) readStdIn() {
	r := bufio.NewReader(os.Stdin)
	cr := newCsvReader(r)
	for {
		// n, err := r.Read(buf[:cap(buf)])
		// buf = buf[:n]
		row, err := cr.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
		if len(row) == 0 {
			continue
		}

		switch row[0] {
		case "exit", "quit":
			bt.wg.Done()
		case "help":
			fmt.Println(HELP_TXT)
		default:
			// 	addCsvRow(row)
			fmt.Println("Row:", row[0])
		}
	}
}

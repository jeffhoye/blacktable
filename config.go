package blacktable

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func (bt *BlackTable) AddConfigFile(fileName string) error {
	fmt.Println("AddConfigFile", fileName)
	return bt.AddCsvConfigFile(fileName)
}

func (bt *BlackTable) AddCsvConfigFile(fileName string) error {
	csvFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	cr := newCsvReader(bufio.NewReader(csvFile))
	for {
		fmt.Println("Reading row from config file")
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

		bt.addCsvConfigRow(row)
	}

	return nil
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
			bt.addCsvConfigRow(row)
		}
	}
}

func (bt *BlackTable) addCsvConfigRow(row []string) {
	fmt.Println("addCsvConfigRow", row[0])
	switch row[0] {
	case "comment":
		return
	}
	period := PeriodicTask{
		Name: row[1],
	}
	switch row[0] {
	case "listen":
		fmt.Println("task listen")
		// task := &Listen {

		// }
	case "send":
		task := &NetworkMessage{
			PeriodicTask: period,
			Protocol:     row[5],
		}
		bt.taskChan <- task
	case "echo":
		fmt.Println("task echo")
	}
	fmt.Println("RowZ:", row[0])

}

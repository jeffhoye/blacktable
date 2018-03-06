package blacktable

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func (bt *BlackTable) AddConfigFile(fileName string) error {
	// fmt.Println("AddConfigFile", fileName)
	return bt.AddCsvConfigFile(fileName)
}

func (bt *BlackTable) AddCsvConfigFile(fileName string) error {
	csvFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	cr := newCsvReader(bufio.NewReader(csvFile))
	rowNum := 0
	for {
		rowNum++
		// fmt.Println("Reading row from config file")
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

		bt.addCsvConfigRow(row, rowNum, fileName)
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
	rowNum := 0
	for {
		rowNum++
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
			bt.addCsvConfigRow(row, rowNum, "stdin")
		}
	}
}

func (bt *BlackTable) parseStart(s string) (t time.Time, err error) {
	var duration time.Duration
	duration, err = bt.parseDuration(s)
	if err != nil {
		return
	}
	if duration < 0 {
		return // time zero
	}
	t = time.Now()
	t = t.Add(duration)
	return
}

func (bt *BlackTable) parsePeriod(t string) (duration time.Duration, err error) {
	duration, err = bt.parseDuration(t)
	if err != nil {
		return
	}
	return
}

func (bt *BlackTable) parseDuration(t string) (duration time.Duration, err error) {
	var seconds float64
	seconds, err = strconv.ParseFloat(t, 64)
	if err != nil {
		return
	}
	return time.Duration(seconds) * time.Second, nil
}

func (bt *BlackTable) parseTimes(t string) (int, error) {
	return strconv.Atoi(t)
}

func (bt *BlackTable) addCsvConfigRow(row []string, line int, fileName string) {
	// fmt.Println("addCsvConfigRow", row[0])
	switch row[0] {
	case "comment":
		return
	}
	if row[0][0] == '#' {
		// handle commented out row
		// fmt.Println("found commented out command", row[0])
		return
	}

	if len(row) < 5 {
		log.Println("Error in", fileName, "line", line, " Not enough columns:", len(row))
	}

	start, err := bt.parseStart(row[2])
	if err != nil {
		log.Println("Error in", fileName, "line", line, " Cant parse start", row[2], err)
		return
	}
	period, err := bt.parsePeriod(row[3])
	if err != nil {
		log.Println("Error in", fileName, "line", line, " Cant parse start", row[3], err)
		return
	}

	times, err := bt.parseTimes(row[4])
	if err != nil {
		log.Println("Error in", fileName, "line", line, " Cant parse start", row[4], err)
		return
	}
	baseTask := PeriodicTask{
		Name:   row[1],
		Start:  start,
		Period: period,
		Times:  times,
	}
	switch row[0] {
	case "receive":
		// fmt.Println("task listen")
		if len(row) < 7 {
			log.Println("Error in", fileName, "line", line, " Not enough columns:", len(row))
		}

		task := &ReceiveTask{
			PeriodicTask: baseTask,
			Protocol:     row[5],
			OnIpPort:     row[6],
		}
		bt.taskChan <- task
	case "send":
		if len(row) < 8 {
			log.Println("Error in", fileName, "line", line, " Not enough columns:", len(row))
		}

		task := &SendTask{
			PeriodicTask: baseTask,
			Protocol:     row[5],
			ToIpPort:     row[6],
			Message:      []byte(row[7]),
		}
		bt.taskChan <- task
	case "echo":
		fmt.Println("task echo")
	}
	// fmt.Println("RowZ:", row[0])

}

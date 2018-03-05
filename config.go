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
	return bt.AddCsvConfigFile(fileName)
}

func (bt *BlackTable) AddCsvConfigFile(fileName string) error {
	csvFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	cr := newCsvReader(bufio.NewReader(csvFile))
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
	fmt.Println("RowZ:", row[0])

}

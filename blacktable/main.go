package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jeffhoye/blacktable"
)

func main() {
	bt, err := blacktable.NewBlackTable()
	if err != nil {
		log.Fatal(err)
	}
	args := os.Args[1:]
	for _, configFileName := range args {
		bt.AddCsvConfigFile(configFileName)
	}

	bt.Start()
	fmt.Println("Exiting")
}

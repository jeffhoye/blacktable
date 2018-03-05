package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jeffhoye/blacktable"
)

func main() {
	// fmt.Println("main1")
	bt, err := blacktable.NewBlackTable()
	if err != nil {
		log.Fatal(err)
	}
	args := os.Args[1:]
	// fmt.Println("main2", len(args))
	for _, configFileName := range args {
		err = bt.AddConfigFile(configFileName)
		if err != nil {
			log.Fatal(err)
		}
	}
	// fmt.Println("main3")
	bt.Start()
	fmt.Println("Exiting")
}

package main

import (
	"fmt"
	"log"
)

func main() {
	// args := os.Args[1:]
	// if len(args) > 0 {
	// 	configFileName = args[0]
	// }

	bt, err := NewBlackTable()
	if err != nil {
		log.Fatal(err)
	}
	bt.Run()
	fmt.Println("Exiting")
}

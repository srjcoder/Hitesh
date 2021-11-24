package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	emptyFile, err := os.Create("emptys.go")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)
	// close the created file
	// emptyFile.Close()
	//
	msg := main.function()
	fmt.Println(msg)

}

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Simple ways to create file and write and read of it")

	content := `"This whole code explain how to use create and read and write into particular file,
					 here we need to note one thing that create may be comes under os interface but
					 write func will come under io interface and read in ioutil"`

	// creating file
	file, err := os.Create("./description.txt")

	// common error func
	checkNilErr(err)

	// inserting content into file
	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is: ", length)

	// close the writing file
	defer file.Close()

	// now read the written file
	readFile("./description.txt")

}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	// simplest way to print byte to string
	fmt.Println("written inside file : ", string(data))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}

}

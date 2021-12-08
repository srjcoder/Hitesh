package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Verify bool
	Age    int
}

func main() {
	fmt.Println("Welcome to struct lets get ur details")

	// initialzing struct (class) to variables
	printOuts := User{"Raju", "Raju@go.dev", true, 22}

	fmt.Printf("my details are: %v\n", printOuts)

	// fetching printout with variables
	fmt.Printf("my details clearly: %+v\n", printOuts)

	// particular fetching
	fmt.Printf("My age %v, in details, \n", printOuts.Age)
}

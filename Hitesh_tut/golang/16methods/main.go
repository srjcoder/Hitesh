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

	// methods with pointer help
	printOuts.NewMail()
	fmt.Println("after NewMail method: ", printOuts.Email)

	printOuts.NewAge()
	fmt.Println("after NewAge method: ", printOuts.Age)
}

func (u *User) NewMail() {
	u.Email = "test@dev.go"
	fmt.Println("Email id is: ", u.Email)
}

func (u User) NewAge() {
	u.Age = 90
	fmt.Println("Email id is: ", u.Age)
}

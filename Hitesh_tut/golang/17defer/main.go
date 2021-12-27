package main

import "fmt"

//defer is predefined func which will print next line of defer
// after print next
// it will perform like last in first out

func main() {
	defer fmt.Println("world")
	defer fmt.Println("three")
	defer fmt.Println("two")
	defer fmt.Println("one")
	fmt.Println("hello")

	useDefer()
}

func useDefer() {
	if true {
		for i := 1; i <= 5; i++ {

			defer fmt.Println(i)

		}
	}
}

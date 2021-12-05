package main

import (
	"fmt"
)

func main() {

	// pointer is using to avoid the problem while using copy of variable
	// instead we can use address of var
	// creating empty pointer

	// var ptr *int
	// fmt.Println("empty ptr", ptr)

	// creating ptr with ' & ' sign which mean empersion using to reference
	myNumber := 62

	var ptr = &myNumber

	// now various prints gonna show up
	// output of variable address
	fmt.Println("the actual value: ", ptr)

	// output of variable
	fmt.Println("the actual value: ", *ptr)

	// the ptr variable with operations
	fmt.Println("the actual value with multi: ", *ptr*2)
}

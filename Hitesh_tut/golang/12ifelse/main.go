package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to ifelse")

	var conditions int = 10

	if conditions > 5 {
		fmt.Println("*****")
	}
	if can := 23; can <= 20 {
		fmt.Println("23")
	}
	if 23%2 == 0 {
		fmt.Println("22")
	} else {
		fmt.Println("nope")
	}
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to array")

	// fruits list
	var fruitLists [5]string

	fruitLists[0] = "apple"
	fruitLists[1] = "tomato"
	fruitLists[4] = "peach"

	fmt.Println("Fruits lists is: ", fruitLists, len(fruitLists))

	// veg list
	var vegLists [5]string

	vegLists[1] = "beans"
	vegLists[2] = "carrot"
	vegLists[3] = "beatroot"

	fmt.Println("vegies are: ", vegLists, len(vegLists))

	// now all arrays
	fmt.Println("all lists: ")

}

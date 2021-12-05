package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to slices practice")

	// initializing array in variable
	var fruitsList = []string{"apple", "peach", "strawberry"}

	// type of var
	fmt.Printf("the fruits lists type is %T\n", fruitsList)

	// printing with string package
	fmt.Println("fruits List with strings.pac: ", fruitsList[0], strings.SplitAfter(fruitsList[2], fruitsList[0]))

	// append method in array
	fmt.Println("appended new value: ", append(fruitsList, "mango", "pineapple", "cherry"))

	// fp fruits list with slices
	fmt.Println("printing list with slices", append(fruitsList[:2]))

	// make package
	var course = []string{"java", "python", "c", "ruby", "perl", "golang"}

	fmt.Println("course var without del: ", course)

	index := 2
	// del var of course
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)
}

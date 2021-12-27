package main

import "fmt"

func main() {

	fmt.Println("welcome to functions concept")
	greetings()

	results := adder(2, 3)
	fmt.Println("results of adder method:", results)

	proRes := proAdder(2, 3, 4, 5, 6, 7)
	fmt.Println("results of proAdder:", proRes)

	proRess, message := proAdder2(2, 3, 4, 5, 6, 7)
	fmt.Println("results of proAdder2:", proRess)
	fmt.Println("results of proAdder2:", message)

}

func greetings() {
	fmt.Println("this is how functions can pass")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0

	for x := range values {
		total += values[x]
	}
	return total
}

func proAdder2(values ...int) (int, string) {
	total := 0

	for x := range values {
		total += values[x]
	}
	return total, "nice and easy"
}

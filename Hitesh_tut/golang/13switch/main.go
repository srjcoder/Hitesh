package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This switch case program was used to print dice throwing in ludo
// now -> using random number from math package
// then implant time.pkg to track time then used UnixNano.pkg similar to nano.pkg to track nano sec
// create 6 nums to throw so for that use rand.pkg with Intn.pkg to only genarate postive numbers
// give 6 and call through switch
// now we can see fallthrough.pkg that is for printing required case with next case

func main() {
	fmt.Println("welcome switch and case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move to 3 spot")
		fallthrough
	case 4:
		fmt.Println("you can move to 4 spot")
	case 5:
		fmt.Println("you can move to 5 spot")
	case 6:
		fmt.Println("you can move to 6 spot")
	default:
		fmt.Println("what is that!")

	}
}

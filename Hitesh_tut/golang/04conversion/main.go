package conversion

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// reciveing rating from users

	fmt.Println("Welcome to star cup shop")
	fmt.Println("Rate the service from 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)

}

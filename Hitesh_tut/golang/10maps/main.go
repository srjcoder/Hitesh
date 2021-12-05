package main

import "fmt"

func main() {
	fmt.Println("welcome to map test")

	// initializing map to var
	languages := make(map[string]string)

	languages["st"] = "stupid"
	languages["bt"] = "bitch"
	languages["io"] = "idiot"
	languages["sob"] = "son of bitch"
	languages["di"] = "damn it"
	languages["bs"] = "bull shit"
	languages["wth"] = "wht ta heck"

	fmt.Println("keys: ", languages)

	// printing with key and values
	for keys, values := range languages {
		fmt.Printf("all keys: %v and values: %v\n", keys, values)

	}

}

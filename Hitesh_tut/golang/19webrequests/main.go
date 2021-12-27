package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// this code is about send request to particular website then get responses from that website
// it will give what we request to it
// here we request about website's body of code

const url = "https://lco.dev"

func main(){
	fmt.Println("Webrequest program demonstrated")

	// Here we are going send get request to one website
	response, err:= http.Get(url)

	if err!=nil{
		panic(err)
	}
	fmt.Printf("response type is: %T\n",response)
	// now we should close request once get it
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}

	content:= string(databytes)
	fmt.Println(content)


}
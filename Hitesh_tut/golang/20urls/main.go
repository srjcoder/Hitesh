package main

import (
	"fmt"
	"net/url"
)

// creating fectices url and extract info from it and
// create info and create url from it

const urls string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj45ghb"

func main() {
	fmt.Println("handling urls :")
	fmt.Println(urls)
	results, _ :=url.Parse(urls)

	// lest see what kind of infos we can extract
	// scheme - its based ssl (better clarification is need)
	// fmt.Println(results.Scheme)

	// host like domain.name
	// fmt.Println(results.Host)

	// path - the folders and sub folders of web
	// fmt.Println(results.Path)

	// generated or created path of webpage
	// fmt.Println(results.RawQuery)

	// port of website
	// fmt.Println(results.Port())

	// now query parameters : url's values
	qparams := results.Query()
	fmt.Printf("the type of query params are: %T\n",qparams)

	// to know the value of url
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	// get those at a time
	for _,val := range qparams{
		fmt.Println("params is: ",val)
	}

	// now constructing url
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/HiteshTut",
		RawPath: "user=raju",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://localhost:3000/books?course=go&payment=true"

func main() {
	fmt.Println("Welcome to urls in Go")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)
	// fmt.Println("result is: ", result)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Println(qparams)
	fmt.Printf("Type of qparams %T\n", qparams)
	fmt.Println(qparams["course"])

	for _, val := range qparams {
		fmt.Println(val)
	}

	partsofurl := &url.URL{
		Scheme:  "https",
		Host:    "noob:3000",
		Path:    "/books",
		RawPath: "user=bot",
	}

	anotherurl := partsofurl.String()
	fmt.Println(anotherurl)
	fmt.Printf("Type of anotherurl is: %T\n", anotherurl)
}

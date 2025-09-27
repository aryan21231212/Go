package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://api.thecatapi.com/v1/breeds"

func main() {
	fmt.Println("Web Requests in Go")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("type of response is %T\n", response)

	defer response.Body.Close() // callers's responsibility to close this.

	bytecode, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(bytecode)

	fmt.Println(content)

}

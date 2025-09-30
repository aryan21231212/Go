package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("hello")
	// getrequest()
	postrequest()
}

func getrequest() {
	const myURL = "http://localhost:3000/get"

	// response, _ := http.Get(myURL)

	// fmt.Println("Status code:", response.StatusCode)
	// fmt.Println("Content length is: ", response.ContentLength)

	// defer response.Body.Close()

	// var reponseString strings.Builder
	// content, _ := io.ReadAll(response.Body)

	// byteCount, _ := reponseString.Write(content)
	// fmt.Println("Byte count is: ", byteCount)
	// fmt.Println("Response string is: ", reponseString.String())
}

func postrequest() {
	const myURL = "http://localhost:3000/post"

	responseData := strings.NewReader(`{
		"name": "Aryan",
		"age":"20",
		"year":"SE"
	}`)

	response, _ := http.Post(myURL, "application/json", responseData)

	content, _ := io.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Println("Content", string(content))
}

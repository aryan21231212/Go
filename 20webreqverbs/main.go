package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("hello")
	// getrequest()
	// postrequest()
	postForm()
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

func postForm() {
	const myURL = "http://localhost:3000/postform"

	data := url.Values{}
	data.Add("name", "aryan")
	data.Add("age", "20")
	data.Add("year", "SE")
	data.Add("city", "pune")

	response, err := http.PostForm(myURL, data)

	if err != nil {
		panic(err)
	}

	content, _ := io.ReadAll(response.Body)
	fmt.Println("Content", string(content))
}

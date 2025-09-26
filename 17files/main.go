package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in Go")
	content := "Welcome to noob's world"
	file, err := os.Create("./noob.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	errorHandling(err)
	fmt.Println("length is: ", length)

	defer file.Close()
	readFile("./noob.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	errorHandling(err)
	fmt.Println("File content: ", string(databyte))
}

func errorHandling(err error) {
	if err != nil {
		panic(err)
	}
}

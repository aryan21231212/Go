package main

import "fmt"

func main() {
	fmt.Println("Welcome to Defer in Go")

	// works on Last in first out principle

	defer fmt.Println("Hello")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")

	rev()
	fmt.Println("World")
}

func rev() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

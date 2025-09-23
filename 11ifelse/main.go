package main

import "fmt"

func main() {
	fmt.Println("if else in Go")

	loginCount := 10

	if loginCount < 10 {
		fmt.Println("Regular user")
	} else if loginCount > 10 {
		fmt.Println("Watch out")
	} else {
		fmt.Println("Exact 10 logins")
	}

	// odd even
	if 5%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// short statement
	if num := 9; num < 0 {
		fmt.Println("Negative number")
	} else {
		fmt.Println("Non-negative number")
	}
}

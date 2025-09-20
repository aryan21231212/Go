package main

import "fmt"

func main() {
	fmt.Println("Noob")

	// pointer declaration

	// var ptr *int     // pointer to an integer
	// var noob *string // pointer to a string

	mynumber := 21

	pointer := &mynumber

	fmt.Println("addrss of mynumber is ", pointer)
	fmt.Println("value at pointer is ", *pointer)

	//operations on pointer
	
	*pointer = *pointer * 2
	fmt.Println("value at pointer after change is ", mynumber)

}

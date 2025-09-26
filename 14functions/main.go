package main

import "fmt"

// nesting of function is not allowed

func main() {
	fmt.Println("hello")
	greeter()
	result := addition(4, 5)
	fmt.Println("Result is: ", result)
	result2 := addition2(4, 5, 6, 7, 8, 8)
	fmt.Println("Sum of addition 2 is: ", result2)
}

func greeter() {
	fmt.Println("Welcome to Functions in Go")
}

// addition of two numbers

func addition(num1 int, num2 int) int {
	return num1 + num2
}

// additon of more than two numbers

func addition2(value ...int) int {
	sum := 0
	for i := range value {
		sum += value[i]
	}
	return sum
}

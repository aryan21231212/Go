package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating  for our pizza:")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating, ", input)
	fmt.Printf("Type of input is: %T \n", input)
}

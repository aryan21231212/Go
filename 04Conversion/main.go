package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello Welcome to our pizza shop !")

	fmt.Println("Rate our pizza (1-5):")

	reader := bufio.NewReader(os.Stdin)
	rating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us ", rating)

	// want to add 1 to the rating
	// convert string to int

	number, _ := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	fmt.Println("Rating increased by 1 is ", number+1)
}

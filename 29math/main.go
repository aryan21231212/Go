package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("welcome to math package")

	//using seed
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println("Random number between 0-100 is:", rand.Intn(100))

	//using crypto
	//crypto is more secure than math/rand
	// import "crypto/rand"
	randNum, _ := rand.Int(rand.Reader, big.NewInt(100))
	fmt.Println("Random number between 0-100 is:", randNum)
}

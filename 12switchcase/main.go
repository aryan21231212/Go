package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch case in Go")

	randomNumner := rand.Intn(6) + 1
	fmt.Println("Random number is", randomNumner)

	switch randomNumner {
	case 1:
		fmt.Println("One, you can open")
	case 2:
		fmt.Println("You can move upto two spots")
	case 3:
		fmt.Println("Three, you can move upto three spots")
	case 4:
		fmt.Println("Four, you can move upto four spots")
		fallthrough // it will execute the next case also
	case 5:
		fmt.Println("Five, you can move upto five spots")
	case 6:
		fmt.Println("Six, you can move upto six spots and roll the dice again")
	default:
		fmt.Println("Noob hai bhai!")
	}
}

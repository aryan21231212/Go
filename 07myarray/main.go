package main

import "fmt"

func main() {

	fmt.Println("Array in Go")

	var superHeros [5]string
	superHeros[0] = "ironman"
	superHeros[1] = "spiderman"
	superHeros[2] = "Hulk"
	superHeros[4] = "thor"

	fmt.Println("Super Heros are: ", superHeros)
	fmt.Println("Length of array is: ", len(superHeros))

	// initialize array

	var myarray = [3]int{21, 19, 45}
	fmt.Println("My array is: ", myarray)
}

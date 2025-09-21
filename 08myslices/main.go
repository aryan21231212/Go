package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in go")

	myfruit := []string{"Apple", "Banana", "Grapes"}
	fmt.Println("My fruit list is", myfruit)

	// adding new fruit to the slice
	myfruit = append(myfruit, "orange", "mango")
	fmt.Println("My fruit list is", myfruit)
	fmt.Printf("Type of fruit list %T \n", myfruit)

	//slicing in slice
	fmt.Println(myfruit[:3])

	//creating slice using make
	var mynumbers = make([]int, 4)
	mynumbers[0] = 100
	mynumbers[1] = 50
	mynumbers[2] = 200
	mynumbers[3] = 130

	// mynumbers[4] = 400 // this will give error as the size of slice is 4

	// to avoid this we use append function
	mynumbers = append(mynumbers, 400, 500)
	fmt.Println("My numbers: ", mynumbers)

	sort.Ints(mynumbers)
	fmt.Println(mynumbers)

	// remove a value from slice based on index
	index := 2
	mynumbers = append(mynumbers[:index], mynumbers[index+1:]...)
	fmt.Println("My numbers: ", mynumbers)
}

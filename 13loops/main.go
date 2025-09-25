package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Go")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i, day := range days {
		fmt.Printf("Day %d is %s\n", i+1, day)
	}

	roughvalue := 1

	for roughvalue < 10 {

		if roughvalue == 2 {
			goto noob
		}

		if roughvalue == 5 {
			break
		}

		fmt.Println("Value is: ", roughvalue)
		roughvalue++

	}
noob:
	fmt.Println("Welcome to noob zone")
}

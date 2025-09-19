package main

import "fmt"

const jwt string = "noobcoder"

func main() {
	var username string = "aryan"
	fmt.Println(username)
	fmt.Printf("The type of username is %T \n", username)

	var islogedin bool = true
	fmt.Println(islogedin)
	fmt.Printf("The type of username is %T \n", islogedin)

	var number uint8 = 255
	fmt.Println(number)
	fmt.Printf("The type of username is %T \n", number)

	var floatnumber float32 = 255.88898383999
	fmt.Println(floatnumber)
	fmt.Printf("The type of username is %T \n", floatnumber)

	var bigfloatnumber float64 = 255.88898383999
	fmt.Println(bigfloatnumber)
	fmt.Printf("The type of username is %T \n", bigfloatnumber)

	var randomvalue int
	fmt.Println(randomvalue)
	fmt.Printf("The type of username is %T \n", randomvalue)

	tokenvalue := 30000.00
	fmt.Println(tokenvalue)
	fmt.Printf("The type of username is %T \n", tokenvalue)

	fmt.Println("The value of Jwt is", jwt)
}

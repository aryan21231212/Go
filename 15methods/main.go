package main

import "fmt"

func main() {

	fmt.Println("Struct in Go")
	// no inheritance in go
	// no super or this in go
	// no constructor in go

	Aryan := User{"Aryan", "noob@go.in", true, 20}
	fmt.Println("User is ", Aryan)

	fmt.Printf("User is %+v\n", Aryan)
	fmt.Printf("Name is %v and email is %v\n", Aryan.Name, Aryan.Email)
	Aryan.gettingStatus()
	Aryan.setter()
	fmt.Printf("Name is %v and email is %v\n", Aryan.Name, Aryan.Email)
}

type User struct {
	Name   string
	Email  string
	status bool
	age    int
}

func (u User) gettingStatus() {
	fmt.Println("Status is: ", u.status)
}

// altering email without passing with reference

func (u User) setter() {
	u.Email = "bot@gmail.com" // not change in original Struct
	fmt.Println("New email is: ", u.Email)
}

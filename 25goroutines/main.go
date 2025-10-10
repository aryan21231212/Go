package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("go routine")
	go greeter("noob")
	greeter("bot")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(2 * time.Microsecond)
		fmt.Println(s)
	}
}

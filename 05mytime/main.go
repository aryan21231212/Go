package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()
	fmt.Println(currTime.Format("02-Jan-2006 Mon 15:04:05"))

	currDate := time.Date(2023, time.June, 20, 12, 23, 0, 0, time.UTC)
	fmt.Println(currDate)

	fmt.Println(currDate.Format(" Mon 15:04:05"))
}

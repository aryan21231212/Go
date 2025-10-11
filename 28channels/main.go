package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("welcome to channels")

	wg := &sync.WaitGroup{}

	mychan := make(chan int, 2)

	wg.Add(2)
	go func(wg *sync.WaitGroup, ch <-chan int) {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		wg.Done()
	}(wg, mychan)
	go func(wg *sync.WaitGroup, ch chan<- int) {
		ch <- 5
		ch <- 6
		close(ch)
		wg.Done()
	}(wg, mychan)

	wg.Wait()
}

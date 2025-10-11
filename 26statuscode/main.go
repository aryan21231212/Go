package main

import (
	"fmt"
	"net/http"
	"sync"
)

var link = []string{"test"}

var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	websites := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
		"https://www.medium.com",
	}

	for _, url := range websites {
		go getStatusCode(url)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("Links:", link)
}

func getStatusCode(url string) {
	defer wg.Done()
	res, _ := http.Get(url)
	println(url, res.StatusCode)
	mut.Lock()
	link = append(link, url)
	mut.Unlock()
}

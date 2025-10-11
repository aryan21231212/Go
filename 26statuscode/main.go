package main

import (
	"net/http"
	"sync"
)

var wg sync.WaitGroup

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
}

func getStatusCode(url string) {
	defer wg.Done()
	res, _ := http.Get(url)
	println(url, res.StatusCode)
}

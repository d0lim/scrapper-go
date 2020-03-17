package main

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

// func main() {
// 	results := map[string]string{}
// 	c := make(chan requestResult)
// 	urls := []string{
// 		"https://www.airbnb.com/",
// 		"https://www.google.com/",
// 		"https://www.amazon.com/",
// 		"https://www.reddit.com/",
// 		"https://soundcloud.com/",
// 		"https://www.facebook.com/",
// 		"https://www.instagram.com/",
// 		"https://academy.nomadcoders.co/",
// 	}
// 	for _, url := range urls {
// 		go hitURL(url, c)
// 	}
// 	for i := 0; i < len(urls); i++ {
// 		result := <-c
// 		results[result.url] = result.status
// 	}
// 	for url, status := range results {
// 		fmt.Println(url, status)
// 	}
// }

func main() {
	results := map[string]string{}
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	c := hitURLChannel(urls...)
	for result := range c {
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}

}

func hitURLChannel(urls ...string) <-chan requestResult {
	out := make(chan requestResult)
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(out chan<- requestResult, url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			status := "OK"
			if err != nil || resp.StatusCode >= 400 {
				status = "FAILED"
			}
			out <- requestResult{url: url, status: status}
		}(out, url)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

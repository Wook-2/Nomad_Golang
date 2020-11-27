package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan result)
	urls := []string{
		"https://www.naver.com",
		"https://www.google.com",
		"https://www.facebook.com",
		"https://github.com/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		r := <-c
		results[r.url] = r.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}

}

var errRequestFailed = errors.New("Request Failed")

func hitURL(url string, c chan<- result) {

	res, err := http.Get(url)

	if err != nil || res.StatusCode >= 400 {
		c <- result{url: url, status: "FAILED"}
	} else {
		c <- result{url: url, status: "OK"}
	}

}

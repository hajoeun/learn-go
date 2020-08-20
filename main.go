package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://www.naver.com",
		"https://www.youtube.com",
		"https://www.banksalad.com",
		"https://www.facebook.com",
		"https://www.github.com",
		"https://www.google.com",
		"https://www.brunch.co.kr",
	}
	results := make(map[string]string)

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

var errRequestFailed = errors.New("Request Failed")

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}

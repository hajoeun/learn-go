package urlchecker

import (
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status int
}

// Urlchecker is URL checker
func Urlchecker() {
	channel := make(chan result)
	urls := []string{
		"https://www.naver.com",
		"https://www.youtube.com",
		"https://www.banksalad.com",
		"https://www.facebook.com",
		"https://www.github.com",
		"https://www.google.com",
		"https://www.brunch.co.kr",
	}

	for _, url := range urls {
		go hitURL(url, channel)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println("Wating for " + urls[i])
		fmt.Println(<-channel)
	}
}

func hitURL(url string, channel chan result) {
	fmt.Println("Checking:", url)
	resp, _ := http.Get(url)

	channel <- result{url: url, status: resp.StatusCode}
}

package urlchecker

import (
	"fmt"
	"net/http"

	"github.com/joeunha/learn-go/urlchecker/types"
)

// Fastchecker testing channel and goroutine
func Fastchecker() {
	channel := make(chan types.Response)
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
		go hitURLFast(url, channel)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-channel)
	}
}

func hitURLFast(url string, channel chan types.Response) {
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		channel <- types.GenResponse(url, "FAILED")
	} else {
		channel <- types.GenResponse(url, "SUCCESS")
	}

}

package nomadcoder

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/joeunha/learn-go/nomadcoder/types"
)

// SlowChecker testing normal function
func SlowChecker() {

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
		fmt.Println(hitURLSlow(url))
	}
}

var errRequestFailed = errors.New("Request Failed")

func hitURLSlow(url string) types.Response {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return types.GenResponse(url, "FAILED")
	}
	return types.GenResponse(url, "SUCCESS")

}

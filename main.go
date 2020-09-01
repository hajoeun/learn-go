package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=javascript&limit=50"

func main() {
	totalPages := getTotalPages()

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println(pageURL)
}

func getTotalPages() int {
	pages := 0
	res, err := http.Get(baseURL)

	checkErr(err)
	checkStatusCode(res.StatusCode)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination-list").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("li").Length() - 1
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatusCode(code int) {
	if code != 200 {
		log.Fatalln("Request failed with Status:", code)
	}
}

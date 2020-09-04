package nomadcoder

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type jobCardData struct {
	id       string
	title    string
	company  string
	location string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=javascript&limit=50"

// JobScrapper scrap job information
func JobScrapper() {
	totalPages := getTotalPages()

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkStatusCode(res.StatusCode)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	jobCards := doc.Find(".jobsearch-SerpJobCard")

	jobCards.Each(func(i int, card *goquery.Selection) {
		id, _ := card.Attr("data-jk")
		title, _ := card.Find("a.jobtitle").Attr("title")
		company := card.Find(".sjcl>.company").Text()
		location := card.Find(".sjcl>.location").Text()
		summary := card.Find(".summary").Text()
		jobData := jobCardData{
			id:       id,
			title:    title,
			company:  company,
			location: location,
			summary:  summary,
		}

		fmt.Println(jobData)
	})
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

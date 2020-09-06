package nomadcoder

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type jobCardData struct {
	id       string
	title    string
	company  string
	location string
	summary  string
}

// JobScrapper scrap job information
func JobScrapper(keyword string) {
	baseURL := "https://kr.indeed.com/jobs?q=" + keyword + "&limit=50"
	var totalJobs []jobCardData
	ch := make(chan []jobCardData)
	totalPages := getTotalPages(baseURL)

	for i := 0; i < totalPages; i++ {
		go getJobs(baseURL, i, ch)
	}

	for i := 0; i < totalPages; i++ {
		jobs := <-ch
		totalJobs = append(totalJobs, jobs...)
	}

	writeJobs(totalJobs)
	fmt.Println("Done! Scrapped", len(totalJobs), "jobs")
}

func writeJobs(jobs []jobCardData) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	header := []string{"id", "title", "company", "location", "summary"}
	wErr := w.Write(header)
	checkErr(wErr)

	for _, job := range jobs {
		job := []string{job.id, job.title, job.company, job.location, job.summary}
		jwErr := w.Write(job)
		checkErr(jwErr)
	}
}

func getJobs(baseURL string, page int, ch chan []jobCardData) {
	var jobs []jobCardData
	c := make(chan jobCardData)
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)

	fmt.Println("Requesting:", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkStatusCode(res.StatusCode)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	jobCards := doc.Find(".jobsearch-SerpJobCard")

	jobCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < jobCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	ch <- jobs
}

func extractJob(card *goquery.Selection, c chan jobCardData) {
	id, _ := card.Attr("data-jk")
	title, _ := card.Find("a.jobtitle").Attr("title")
	company := CleanString(card.Find(".sjcl .company").Text())
	location := CleanString(card.Find(".sjcl .location").Text())
	summary := CleanString(card.Find(".summary").Text())
	c <- jobCardData{
		id:       id,
		title:    title,
		company:  company,
		location: location,
		summary:  summary,
	}
}

// CleanString is clean string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getTotalPages(baseURL string) int {
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

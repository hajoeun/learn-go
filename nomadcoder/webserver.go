package nomadcoder

import (
	"os"
	"strings"

	"github.com/joeunha/learn-go/nomadcoder"

	"github.com/labstack/echo"
)

func WebServer() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.File("./nomadcoder/index.html")
	})

	e.POST("/scrape", func(c echo.Context) error {
		defer os.Remove("jobs.csv")
		keyword := strings.ToLower(nomadcoder.CleanString(c.FormValue("keyword")))
		nomadcoder.JobScrapper(keyword)
		return c.Attachment("jobs.csv", "jobs.csv")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

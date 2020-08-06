package main

import (
	"os"
	"strings"

	"github.com/labstack/echo"

	"github.com/vapvin/Go/scrapper"
)

const file_name string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(file_name)
	term := strings.ToLower(scrapper.CleanStr(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(file_name, file_name)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}

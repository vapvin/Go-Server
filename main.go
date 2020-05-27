package main

import (
	"strings"

	"github.com/labstack/echo"

	"github.com/vapvin/Go/scrapper"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	term := strings.ToLower(scrapper.CleanStr(c.FormValue("term")))
	return c.File(term)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}

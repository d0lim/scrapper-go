package main

import (
	"os"
	"strings"

	"github.com/imdigo/scrapper-go/scrapper"
	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("index.html")
}

func handleScrap(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrap(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrap", handleScrap)
	e.Logger.Fatal(e.Start(":1323"))

}

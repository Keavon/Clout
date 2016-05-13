package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := echo.New()

	// Serve index page
	e.File("/", "index.html")

	// Serve static assets from files
	e.Static("/css", "css")
	e.Static("/js", "js")
	e.Static("/icons", "icons")

	e.Run(standard.New(":" + os.Getenv("PORT")))
}

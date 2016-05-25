package main

import (
	"os"

	"github.com/keavon/clout/pkg/api"
	"github.com/keavon/clout/pkg/pool"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	rp := pool.NewPool(os.Getenv("REDIS_URL"))

	a := api.New(rp)

	e := echo.New()

	// Uncomment to start debug mode
	//e.SetDebug(true)

	// Configure middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Serve index page
	e.File("/", "index.html")

	e.POST("/api/create", a.Create)
	e.POST("/api/join", a.Join)
	e.GET("/api/player", a.Player, a.Auth())
	e.POST("/api/purchase/:id", a.Purchase, a.Auth())
	e.POST("/api/operational/:id", a.Operational, a.Auth())
	e.POST("/api/start", a.Start, a.Auth())
	e.GET("/api/rankings/:game", a.Rankings)

	// Serve static assets from files
	e.Static("/css", "css")
	e.Static("/js", "js")
	e.Static("/icons", "icons")

	e.Run(standard.New(":" + os.Getenv("PORT")))
}

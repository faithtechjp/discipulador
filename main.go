package main

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"time"

	"github.com/hackformissions/discipulador/model"
	"github.com/hackformissions/discipulador/storage"
)

var db *storage.Store

func main() {
	var err error

	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// storage
	db, err = storage.Init("./data")
	if err != nil {
		log.Panicf("Failure. err=%s", err)
	}

	// Echo instance
	e := echo.New()
	e.Static("/", "static")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Set Renderer
	e.Renderer = echoview.Default()

	// Routes
	e.GET("/", func(c echo.Context) error {
		//render with master
		return c.Render(http.StatusOK, "index", echo.Map{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	e.GET("/page", func(c echo.Context) error {
		//render only file, must full name with extension
		return c.Render(http.StatusOK, "page.html", echo.Map{"title": "Page file title!!"})
	})

	// Start server
	e.Logger.Fatal(e.Start(":9090"))
}

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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
			"title": "Preciosos",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	e.GET("/discipulado", func(c echo.Context) error {
		//render with master
		db.Mu.RLock()
		defer db.Mu.RUnlock()
		pp := db.UnsafeReadAllPersons()
		return c.Render(http.StatusOK, "discipulado", echo.Map{
			"title":   "Preciosos: Lista de Discipulados cadastrados",
			"persons": pp,
			"now":     time.Now().Format("Mon Jan _2 15:04:05 2006"),
		})
	})

	// Routes: discipulados (REST API)
	e.GET("/discipulados", func(c echo.Context) error {
		db.Mu.RLock()
		defer db.Mu.RUnlock()
		pp := db.UnsafeReadAllPersons()
		return c.JSON(http.StatusOK, pp)
	})
	e.GET("/discipulados/:id", func(c echo.Context) error {
		db.Mu.RLock()
		defer db.Mu.RUnlock()
		pp := db.UnsafeReadAllPersons()
		return c.JSON(http.StatusOK, pp[c.Param("id")])
	})
	e.DELETE("/discipulados/:id", func(c echo.Context) error {
		err := db.Delete(storage.PERSON_STORE, c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "FAIL")
		}
		return c.JSON(http.StatusOK, "OK")
	})
	e.POST("/discipulados", func(c echo.Context) error {
		p := new(model.Person)
		if err = c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		jsonString, _ := json.Marshal(p)
		hashId := sha256.Sum256([]byte(jsonString))
		p.ID = hex.EncodeToString(hashId[:])
		if err = db.Write(storage.PERSON_STORE, p.ID, p); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, p)
	})

	e.GET("/page", func(c echo.Context) error {
		//render only file, must have full name with extension
		return c.Render(http.StatusOK, "page.html", echo.Map{"title": "Page file title!!"})
	})

	// Start server
	e.Logger.Fatal(e.Start(":9090"))
}

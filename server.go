package main

import (
	"net/http"
	"os"
	"portfolio/templates"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static files
	e.Static("/assets", "assets")
	e.Static("/dist", "dist")

	// init template renderer
	templates.NewTemplateRenderer(e, "public/*.html")

	// routes
	e.GET("/", func(c echo.Context) error {
		gaID := os.Getenv("GA_ID")
		data := struct {
			GAID string
		}{
			GAID: gaID,
		}

		return c.Render(http.StatusOK, "index", data)
	})

	// To use when have blogs
	// e.GET("/redirects", func(c echo.Context) error {
	// 	c.Response().Header().Set("HX-Redirect", "/test")
	// 	return c.NoContent(http.StatusOK)
	// })

	e.Logger.Fatal(e.Start(":1323"))
}
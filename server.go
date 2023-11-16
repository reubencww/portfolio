package main

import (
	"net/http"
	"portfolio/templates"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
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
		return c.Render(http.StatusOK, "index", nil)
	})

	// To use when have blogs
	// e.GET("/redirects", func(c echo.Context) error {
	// 	c.Response().Header().Set("HX-Redirect", "/test")
	// 	return c.NoContent(http.StatusOK)
	// })

	e.Logger.Fatal(e.Start(":1323"))
}
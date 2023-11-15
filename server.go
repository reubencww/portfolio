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

	// routes to be deleted
	e.GET("/hello", func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello", nil)
	})

	e.GET("/test", func(c echo.Context) error {
		return c.Render(http.StatusOK, "test", nil)
	})

	e.POST("/redirects", func(c echo.Context) error {
		c.Response().Header().Set("HX-Redirect", "/test")
		return c.NoContent(http.StatusOK)

	})

	e.GET("/get-info", func(c echo.Context) error {
		res := map[string]interface{}{
			"Name": "Reuben",
			"Github": "github.com/reubencww",
			"Email": "reubenchongww2512@gmail.com",
		}
		return c.Render(http.StatusOK, "name_card", res)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
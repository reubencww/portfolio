package main

import (
	"net/http"
	"portfolio/templates"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// init template renderer
	templates.NewTemplateRenderer(e, "public/*.html")

	// routes
	e.GET("/hello", func(c echo.Context) error {
		res := map[string]interface{} {
			"Name": "Reuben",
		}
		return c.Render(http.StatusOK, "index", res)
	})

	e.GET("/test", func(c echo.Context) error {
		return c.Render(http.StatusOK, "test", nil)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
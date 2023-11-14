package templates

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Template renderer for echo framework
// https://echo.labstack.com/docs/templates
type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func NewTemplateRenderer(e *echo.Echo, paths string) {
	tmpl := template.Must(template.ParseGlob(paths))
	renderer := newTemplate(tmpl)
	e.Renderer = renderer
}

func newTemplate(templates *template.Template) echo.Renderer {
	return &Template {
		Templates: templates,
	}
}
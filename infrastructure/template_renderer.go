package infrastructure

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render ...
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) (err error) {
	err = t.templates.ExecuteTemplate(w, name, data)
	return
}

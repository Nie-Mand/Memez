package config

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

const (
	TEMPLATES_DIR = "templates"
)
type TemplateRenderer struct {
    templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    if viewContext, isMap := data.(map[string]interface{}); isMap {
        viewContext["reverse"] = c.Echo().Reverse
    }
    return t.templates.ExecuteTemplate(w, name, data)
}

func InjectRenderer(e *echo.Echo) {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob(TEMPLATES_DIR + "/*.html")),
	}
	e.Renderer = renderer
}
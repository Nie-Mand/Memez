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

type RendererConfigurator func(*TemplateRenderer) error

func WithTemplatesDir(t string) RendererConfigurator {
	return func(renderer *TemplateRenderer) error {
		renderer.templates = template.Must(template.ParseGlob(t + "/*.html"))
		return nil
	}
}

func InjectRenderer(e *echo.Echo, rc ...RendererConfigurator) {
	renderer := &TemplateRenderer{}

	for _, c := range rc {
		c(renderer)
	}

	if renderer.templates == nil {
		renderer.templates = template.Must(template.ParseGlob(TEMPLATES_DIR + "/*.html"))
	}
	e.Renderer = renderer
}
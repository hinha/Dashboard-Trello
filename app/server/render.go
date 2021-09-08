package server

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
	debug     bool
	location  string
}

func templateRenderer(location string, debug bool) *TemplateRenderer {
	//stat, _ := os.Getwd()
	return &TemplateRenderer{
		location:  location,
		debug:     debug,
		templates: template.Must(template.ParseGlob(location)),
	}
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// hot reload mode
	if t.debug {
		t.templates = template.Must(template.ParseGlob(t.location))
	}

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

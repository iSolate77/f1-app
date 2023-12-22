package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TemplateReader struct {
	templates *template.Template
}

func (t *TemplateReader) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	tmpls, err := template.ParseGlob("web/templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %s", err.Error())
	}

	e := echo.New()
	e.Static("/static", "web/static")

	e.Renderer = &TemplateReader{
		templates: tmpls,
	}

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "base.html", nil)
	})

	log.Println("Server started on port 8080")
	e.Start(":8080")
}

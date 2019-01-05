package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ------
// types
// ------

// TemplateRegistry ...
type TemplateRegistry struct {
	Templates map[string]*template.Template
}

// ------------
// entry point
// ------------

func main() {
	// Echo instance
	e := echo.New()

	// Debug mode
	e.Debug = true

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	// Static assets
	e.Static("/static", "asset")

	// Templates
	templates := make(map[string]*template.Template)
	templates["hello"] = registerTemplate("layout", "hello")
	templates["howdy"] = registerTemplate("layout", "howdy")

	// Renerer
	e.Renderer = &TemplateRegistry{Templates: templates}

	// Route => handler
	e.GET("/", helloHandler)
	e.GET("/howdy", howdyHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// ---------
// handlers
// ---------
func helloHandler(c echo.Context) error {
	data := map[string]interface{}{
		"title": "Hello",
		"now":   time.Now().Format(time.RFC3339),
	}
	return c.Render(http.StatusOK, "hello", data)
}

func howdyHandler(c echo.Context) error {
	data := map[string]interface{}{
		"title": "Howdy",
		"now":   time.Now().Format(time.RFC3339),
	}
	return c.Render(http.StatusOK, "howdy", data)
}

// -------
// custom
// -------

// Render ...
func (tr *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	t, ok := tr.Templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}

	// Add global methods if data is a map
	if value, ok := data.(map[string]interface{}); ok {
		value["reverse"] = c.Echo().Reverse
	}

	return t.ExecuteTemplate(w, "layout", data)
}

// ------
// utils
// ------

func registerTemplate(filenames ...string) *template.Template {
	var files []string
	for _, value := range filenames {
		files = append(files, fmt.Sprintf("template/%s.html", value))
	}

	return template.Must(template.ParseFiles(files...))
}

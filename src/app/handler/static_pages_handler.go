package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// ---------
// handlers
// ---------

// HomeHandler ...
func HomeHandler(c echo.Context) error {
	data := map[string]interface{}{
		"title": "Hello",
		"now":   time.Now().Format(time.RFC3339),
	}
	return c.Render(http.StatusOK, "home", data)
}

// HelpHandler ...
func HelpHandler(c echo.Context) error {
	data := map[string]interface{}{
		"title": "Hello",
		"now":   time.Now().Format(time.RFC3339),
	}
	return c.Render(http.StatusOK, "help", data)
}

// AboutHandler ...
func AboutHandler(c echo.Context) error {
	data := map[string]interface{}{
		"title": "Hello",
		"now":   time.Now().Format(time.RFC3339),
	}
	return c.Render(http.StatusOK, "about", data)
}

// ContactHandler ...
func ContactHandler(c echo.Context) error {
	data := map[string]interface{}{
		"title": "Hello",
		"now":   time.Now().Format(time.RFC3339),
	}
	return c.Render(http.StatusOK, "contact", data)
}

// HowdyHandler ...
func HowdyHandler(c echo.Context) error {
	data := map[string]interface{}{
		"title": "Howdy",
		"now":   time.Now().Format(time.RFC3339),
	}
	return c.Render(http.StatusOK, "howdy", data)
}

// ParrotHandler ...
func ParrotHandler(c echo.Context) error {
	message := c.Param("message")
	return c.String(http.StatusOK, message)
}

package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// ---------
// handlers
// ---------

// HelloHandler ...
func HelloHandler(c echo.Context) error {
	data := map[string]interface{}{
		"title": "Hello",
		"now":   time.Now().Format(time.RFC3339),
	}
	return c.Render(http.StatusOK, "hello", data)
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

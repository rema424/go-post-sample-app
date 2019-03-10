package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// ---------
// handlers
// ---------

// UsersNewHandler ...
func UsersNewHandler(c echo.Context) error {
	data := map[string]interface{}{
		"title": "Hello",
		"now":   time.Now().Format(time.RFC3339),
	}
	return c.Render(http.StatusOK, "users/new", data)
}

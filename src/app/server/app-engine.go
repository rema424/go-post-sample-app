// +build appengine

package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func createMux() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is app-engine.go createMux.")
	})

	// note: we don't need to provide the middleware or static handlers, that's taken care of by the platform
	// app engine has it's own "main" wrapper - we just need to hook echo into the default handler
	http.Handle("/", e)
	return e
}

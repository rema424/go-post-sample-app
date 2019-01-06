package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	// hook into the echo instance to create an endpoint group
	// and add specific middleware to it plus handlers
	g := e.Group("/")
	g.Use(middleware.CORS())

	g.GET("", helloHandler)
	g.GET("howdy", howdyHandler)
	g.GET(":message", parrotHandler)
}

func helloHandler(c echo.Context) error {
	return c.String(http.StatusCreated, "Hello, World!")
}

func howdyHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Howdy, World")
}

func parrotHandler(c echo.Context) error {
	message := c.Param("message")
	return c.String(http.StatusOK, message)
}

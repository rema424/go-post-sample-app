// +build !appengine,!appenginevm

package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func createMux() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is app-standalone.go createMux.")
	})

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Static("/", "public")

	return e
}

func main() {
	fmt.Println("This is app-standalone.go's main.")
	e.Logger.Fatal(e.Start(":8080"))
}

// +build !appengine,!appenginevm

package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func createMux() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	return e
}

func main() {
	e.Logger.Fatal(e.Start(":8080"))
}

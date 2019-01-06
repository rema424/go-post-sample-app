// +build appenginevm

package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"google.golang.org/appengine"
)

func createMux() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is app-standalone.go createMux.")
	})
	// note: we don't need to provide the middleware or static handlers
	// for the appengine vm version - that's taken care of by the platform
	return e
}

func main() {
	fmt.Println("This is app-managed.go's main.")
	// the appengine package provides a convenient method to handle the health-check requests
	// and also run the app on the correct port. We just need to add Echo to the default handler
	e := echo.New(":8080")
	http.Handle("/", e)
	appengine.Main()
}

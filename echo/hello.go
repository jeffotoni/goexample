/*
* Example framework echo
*
* @package     main
* @author      @jeffotoni
* @size        18/07/2017
*
 */

package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Hello(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, World!\n")
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", Hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

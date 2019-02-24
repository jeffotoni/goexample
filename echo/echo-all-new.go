// Go Api server
// @jeffotoni
// 2019-02-22

package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// TYPES START POSSIBLE
// App é uma instancia de echo
//var App = echo.New()
//var App echo.Echo
// var App *echo.Echo
var App = echo.Echo{}

func init() {

	App := echo.New()

	//Instancia o echo
	//App.echo.New()

	//Middleware
	App.Use(middleware.Logger())
	App.Use(middleware.Recover())

	//Rota
	App.GET("/", home)
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func main() {

	fmt.Println("ok")
}

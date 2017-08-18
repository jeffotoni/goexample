/*
* Example sync.Mutex
*
* @package     main
* @author      @jeffotoni
* @size        18/08/2017
*
 */

package main

import (
	"github.com/gin-gonic/gin"
	"net/http3"
)

func main() {

	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}

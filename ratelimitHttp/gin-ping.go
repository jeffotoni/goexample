/*
* Example tollbooth width gin-gonic
*
* @package     main
* @author      @jeffotoni
* @size        15/07/2017
*
 */

package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:8081") // listen and serve on 0.0.0.0:8080
}

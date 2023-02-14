/*
* Example sync.Mutex
*
* @package     main
* @author      @jeffotoni
* @size        18/08/2017
*
 */

package main

import "github.com/gin-gonic/gin"

func loginEndpoint(c *gin.Context) {

	c.String(200, "Success")
}

func submitEndpoint(c *gin.Context) {

	c.String(200, "Success")
}

func readEndpoint(c *gin.Context) {

	c.String(200, "Success")
}

func main() {

	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}

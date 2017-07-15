/*
* Example tollbooth width gin-gonic
*
* @package     main
* @author      @jeffotoni
* @size        15/07/2017
*
 */

package main

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/thirdparty/tollbooth_gin"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

	gin.SetMode(gin.DebugMode)

	r := gin.New()

	// Create a limiter
	limiter := tollbooth.NewLimiter(2, time.Second)

	limiter = tollbooth.NewLimiterExpiringBuckets(5, time.Second, time.Hour, 0)

	limiter.IPLookups = []string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"}

	limiter.Methods = []string{"GET", "POST"}

	r.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{

			"msg": "pong",
		})
	})

	//
	//
	//
	r.GET("/login", tollbooth_gin.LimitHandler(limiter), func(c *gin.Context) {

		c.String(200, `{"msg":"ok, token acess"}`)
	})

	r.Run(":12345")
}

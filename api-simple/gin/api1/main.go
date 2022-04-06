package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// func middleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		body := c.Request.Body
// 		x, _ := ioutil.ReadAll(body)
// 		fmt.Printf("%s \n", string(x))
// 		fmt.Println("I am a middleware for json schema validation")
// 		c.Next()
// 		return
// 	}
// }

type UserS struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	v1 := r.Group("/v1")

	//router.Use(middleware())
	v1.POST("/user", User)
	v1.POST("/user2", User2)
	v1.POST("/user3", User3)
	v1.POST("/user4", User4)
	v1.POST("/user5", User5)

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//Listen and serve
	log.Println("Run Server: 8080")
	r.Run("0.0.0.0:8080")
}

func jsonHeader(c *gin.Context) (strjson string) {
	var m = make(map[string]string)
	for k, v := range c.Request.Header {
		m[k] = strings.Join(v, "")
	}
	b, err := json.Marshal(&m)
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{"code": 404, "message": "Error Json Marshal"})
		return
	}
	strjson = string(b)
	return
}
func User5(c *gin.Context) {
	var m = make(map[string]string)
	for k, v := range c.Request.Header {
		m[k] = strings.Join(v, "")
	}
	b, err := json.Marshal(&m)
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{"code": 404, "message": "Error Json Marshal"})
		return
	}
	c.Writer.Header().Set("MessageId", "X39393X3936393")
	c.String(200, string(b))
}

func User4(c *gin.Context) {
	jsonHeader := jsonHeader(c)
	log.Println(jsonHeader)

	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	jsonBody := buf.Bytes()
	log.Println(string(jsonBody))

	var u UserS
	err := json.Unmarshal(jsonBody, &u)
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{"code": 404, "message": "Error Json Unmarshal"})
		return
	}
	c.Writer.Header().Set("MessageId", "X39393X3936393")
	c.JSON(http.StatusOK, u)
}

func User3(c *gin.Context) {
	jsonBody, err := c.GetRawData()
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{"code": 404, "message": "Error Json Payload"})
		return
	}

	if len(jsonBody) == 0 {
		c.JSON(400, gin.H{"code": 400, "message": "Body empty"})
		return
	}

	log.Println(string(jsonBody))

	var u UserS
	err = json.Unmarshal(jsonBody, &u)
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{"code": 404, "message": "Error Json Unmarshal"})
		return
	}
	c.Writer.Header().Set("MessageId", "X39393X3936393")
	c.JSON(200, u)
}

func User2(c *gin.Context) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	jsonBody := buf.String()
	log.Println(string(jsonBody))
	c.String(http.StatusOK, jsonBody)
}

func User(c *gin.Context) {
	var u UserS
	if err := c.ShouldBindJSON(&u); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"code": 400, "message": "Json with error"})
		return
	}
	c.Writer.Header().Set("MessageId", "X39393X3936393")
	c.IndentedJSON(http.StatusOK, u)
	// c.JSON(http.StatusOK, u)
}

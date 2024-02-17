package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type MyStruct struct {
	Codigo string `json:"codigo"`
}

func (obj MyStruct) myHandler(c *gin.Context) {
	var jsonData map[string]interface{}

	err := c.BindJSON(&jsonData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jsonData)
}

func main() {
	router := gin.Default()
	obj := MyStruct{
		Codigo:"98383",
	}
	router.POST("/endpoint", obj.myHandler)
	fmt.Println("Run Server :8080")
	router.Run(":8080")
}


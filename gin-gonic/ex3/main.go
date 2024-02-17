package main

import (
    "github.com/gin-gonic/gin"
)

type myStruct struct {
    codigo string
}

func (obj *myStruct) myHandler(c *gin.Context) {
    var reqBody map[string]interface{}
    
    if err := c.ShouldBindJSON(&reqBody); err != nil {
	    c.JSON(400, gin.H{"error":err.Error()})
        return
    }

    c.JSON(200, reqBody)
}

func main() {
    r := gin.Default()
    
    obj := &myStruct{
        codigo: "12345",
    }
    r.POST("/endpoint/:codigo", obj.myHandler)
    
    r.Run()
}

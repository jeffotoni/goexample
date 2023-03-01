package main

import (
	"fmt"

	"github.com/gojeffotoni/quick"
)

// curl -i -XPOST -H "Content-Type:application/json" 'localhost:8080/v1/user' -d '{"name":"marcus", "year":1982}'
func main() {

	app := quick.New()
	app.Post("/v1/user", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")
		type My struct {
			Name string `json:"name"`
			Year int    `json:"year"`
		}

		var my My
		err := c.BodyParser(&my)
		if err != nil {
			c.Status(400).SendString(err.Error())
			return
		}

		fmt.Println("String:", c.BodyString())
		fmt.Println("byte:", c.Body())
		c.Status(200).JSON(&my)
		return
	})

	app.Listen("0.0.0.0:8080")
}

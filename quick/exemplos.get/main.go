package main

import "github.com/gojeffotoni/quick"

// curl -i -XGET localhost:8080/v1/user
func main() {
	app := quick.New()

	app.Get("/v1/user", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")
		c.Status(200).SendString("Quick em ação ❤️!")
	})

	app.Listen("0.0.0.0:8080")
}

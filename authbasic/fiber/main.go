package main

import (
	"github.com/gofiber/fiber/v2"
	 "github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app := fiber.New()

	auth := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "password",
		},
	})

	app.Use(auth)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, you've been authenticated!")
	})
	app.Listen(":8080")
}


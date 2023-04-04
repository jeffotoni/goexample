package main

import (
	"github.com/jeffotoni/quick"
)

func main() {
	q := quick.New()

	q.Get("/", func(c *quick.Ctx) error {
		return c.Status(quick.StatusOK).SendString("Bem-vindo ao meu site!")
	})

	q.Get("/api", func(c *quick.Ctx) error {
		return c.Status(quick.StatusOK).SendString("Você acessou a API!")
	})

	q.Listen(":8080")
}

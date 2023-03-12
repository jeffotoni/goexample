package handler

import (
	"github.com/gofiber/fiber"
	"github.com/jeffotoni/goexample/api/fiber/crypt"
)

func Check(c *fiber.Ctx) error {
	ID, err := crypt.RandomID()
	if err != nil {
		return c.Status(400).SendString(`{"msg":"Bad Request RandomID"}`)
	}
	c.Set("ID", ID)
	return c.Status(200).SendString(`{"msg":"ok", "token":"xxxxxxxxxxxxxxxxxx3993"}`)
}

func Healthz(c *fiber.Ctx) error {
	ID, err := crypt.RandomID()
	if err != nil {
		return c.Status(400).SendString(`{"msg":"Bad Request RandomID"}`)
	}

	c.Set("ID", ID)
	c.Set("Content-Type", "application/json")
	//return c.Status(200)
	return c.Status(200).SendString("")
}

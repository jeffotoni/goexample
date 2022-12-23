package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit:                5 * 1024,
		DisableStartupMessage:    false,
		Prefork:                  false,
		Concurrency:              1 * 1024 * 1024,
		DisableHeaderNormalizing: false,
		//DisableKeepalive: true,
	})

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		Output:     os.Stdout,
	}))
	app.Get("/healthz", Healthz)

	log.Fatal(app.ListenTLS("0.0.0.0:8080", "cert.pem", "key.pem"))
}

func Healthz(c *fiber.Ctx) error {
	return c.Status(200).SendString("")
}

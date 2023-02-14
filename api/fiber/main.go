// Go Api server
// @jeffotoni
package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jeffotoni/goexample/api/fiber/handler"
)

func main() {
	app := fiber.New(
		fiber.Config{
			BodyLimit:                1 * 512,
			DisableStartupMessage:    false,
			Prefork:                  false,
			Concurrency:              1 * 1024 * 1024,
			DisableHeaderNormalizing: false,
		},
	)
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		Output:     os.Stdout,
	}))

	app.Get("/healthz", handler.Healthz)
	app.Post("/auth/check", handler.Check)
	app.Listen("0.0.0.0:8080")
}

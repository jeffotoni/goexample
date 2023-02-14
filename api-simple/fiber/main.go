package main

import "github.com/gofiber/fiber/v2"

type userPost struct {
  Name string `json:"name"`
  Cpf string  `json:"cpf"`
  Year int `json:"year"`
}

func main() {
	app := fiber.New()
  app.Get("/api/v1/ping", Ping)
  app.Post("/api/v1/user", User)
	app.Listen(":8080")
}

func Ping(c *fiber.Ctx) error {
  c.Set("Content-Type", "application/json")
  return c.Status(200).SendString(`{"msg":"pong"}`)
}

func User(c *fiber.Ctx) (err error) {
  var user userPost
  err = c.BodyParser(&user)
  if err !=nil{
    return c.Status(400).SendString(`{"msg":"error parse, json:`+err.Error()+`"}`)
  }
  return c.Status(200).JSON(user)
}

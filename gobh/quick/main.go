package main 

import "github.com/jeffotoni/quick"

func main() {

	q := quick.New()

	q.Get("/v1/user", func(c *quick.Ctx) error{
		c.Set("Content-Type", "application/json")
		return c.Status(200).String("Evento Go ❤️")
	})

	q.Listen(":8080")
}









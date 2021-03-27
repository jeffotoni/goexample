package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {

		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)

		// c.Locals is added to the *websocket.Conn
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			//mt  int
			msg []byte
			err error
		)
		for {
			println("new client:", c.Params("id"))
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}

			fmt.Sprintf("%s", msg)
			//log.Printf("recv: %s", msg)

			// if err = c.WriteMessage(mt, msg); err != nil {
			// 	log.Println("write:", err)
			// 	break
			// }

			ticker := time.NewTicker(time.Millisecond * 100)
			defer ticker.Stop()

			for {
				select {

				case t := <-ticker.C:
					//fmt.Sprintf("%s", t.String())
					jsonStr := `{"nome":"jefferson","code":"user_` + c.Params("id") + `_9939393xxxx", "time":"` + t.String() + `"}`
					err := c.WriteMessage(websocket.TextMessage, []byte(jsonStr))
					if err != nil {
						log.Println("write:", err)
						break
					}
				case <-interrupt:
					log.Println("interrupt")
					// Cleanly close the connection by sending a close message and then
					// waiting (with timeout) for the server to close the connection.
					err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
					if err != nil {
						log.Println("write close:", err)
						break
					}
					select {
					case <-time.After(time.Millisecond):
					}
					os.Exit(0)
				}
			}
		}

	}))

	log.Fatal(app.Listen(":3000"))
	// Access the websocket server: ws://localhost:3000/ws/123?v=1.0
	// https://www.websocket.org/echo.html
}

/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

var err error

func WriteScren() {

	fmt.Println("Hello, new method client....")
	time.Sleep(time.Millisecond * 300)
}

func ListenWebSocker(ws *websocket.Conn) {

	for { // loop

		var reply string //  Receive the message by reference

		// Receiving client message
		websocket.Message.Receive(ws, &reply)
		if len(reply) > 0 {
			msg := `{"name":"jeffotoni", "code":"x393993993_` + reply + `"}`
			websocket.Message.Send(ws, msg) // Sending message to the client

			//WriteScren()
			time.Sleep(time.Millisecond * 300)
			//fmt.Println("Client sent: " + reply)
			//println("send...")
		}
	}
}

func main() {

	// Defining our api
	http.Handle("/chat", websocket.Handler(ListenWebSocker))

	// Opening the door for receiving messages
	if err := http.ListenAndServe(":1234", nil); err != nil {

		log.Fatal("ListenAndServe:", err)

	}

}

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

	"golang.org/x/net/websocket"
)

var err error

func ListenWebSocker(ws *websocket.Conn) {

	for { // loop

		var reply string //  Receive the message by reference

		// Receiving client message
		websocket.Message.Receive(ws, &reply)

		fmt.Println("Client sent: " + reply)

		msg := "Hello I'm the websocket server!"
		websocket.Message.Send(ws, msg) // Sending message to the client
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

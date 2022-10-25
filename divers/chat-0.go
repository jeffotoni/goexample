/*
	A minimal Go TCP chat example.
	Run this like
		> go run chat-0.go
	That will run a TCP chat server at localhost:6000.
	You can connect to that chat server like
		> telnet localhost 6000
	And, of course, others can connect using your IP
	address like
		> telnet YOUR-IP-HERE 6000
	assuming your firewall allows it.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	// Number of people whom ever connected
	//
	clientCount := 0

	// All people who are connected; a map wherein
	// the keys are net.Conn objects and the values
	// are client "ids", an integer.
	//
	allClients := make(map[net.Conn]int)

	// Channel into which the TCP server will push
	// new connections.
	//
	newConnections := make(chan net.Conn)

	// Channel into which we'll push dead connections
	// for removal from allClients.
	//
	deadConnections := make(chan net.Conn)

	// Channel into which we'll push messages from
	// connected clients so that we can broadcast them
	// to every connection in allClients.
	//
	messages := make(chan string)

	// Start the TCP server
	//
	server, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Tell the server to accept connections forever
	// and push new connections into the newConnections channel.
	//
	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			newConnections <- conn
		}
	}()

	// Loop endlessly
	//
	for {

		// Handle 1) new connections; 2) dead connections;
		// and, 3) broadcast messages.
		//
		select {

		// Accept new clients
		//
		case conn := <-newConnections:

			log.Printf("Accepted new client, #%d", clientCount)

			// Add this connection to the `allClients` map
			//
			allClients[conn] = clientCount
			clientCount += 1

			// Constantly read incoming messages from this
			// client in a goroutine and push those onto
			// the messages channel for broadcast to others.
			//
			go func(conn net.Conn, clientId int) {
				reader := bufio.NewReader(conn)
				for {
					incoming, err := reader.ReadString('\n')
					if err != nil {
						break
					}
					messages <- fmt.Sprintf("Client %d > %s", clientId, incoming)
				}

				// When we encouter `err` reading, send this
				// connection to `deadConnections` for removal.
				//
				deadConnections <- conn

			}(conn, allClients[conn])

		// Accept messages from connected clients
		//
		case message := <-messages:

			// Loop over all connected clients
			//
			for conn, _ := range allClients {

				// Send them a message in a go-routine
				// so that the network operation doesn't block
				//
				go func(conn net.Conn, message string) {
					_, err := conn.Write([]byte(message))

					// If there was an error communicating
					// with them, the connection is dead.
					if err != nil {
						deadConnections <- conn
					}
				}(conn, message)
			}
			log.Printf("New message: %s", message)
			log.Printf("Broadcast to %d clients", len(allClients))

		// Remove dead clients
		//
		case conn := <-deadConnections:
			log.Printf("Client %d disconnected", allClients[conn])
			delete(allClients, conn)
		}
	}
}

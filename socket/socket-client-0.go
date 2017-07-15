/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"io"
	"log"
	"net"
	"time"
)

var (
	SOCK = "/tmp/unix.sock"
)

func readerServer(r io.Reader) {

	// Creating the buffer
	buf := make([]byte, 1024)

	// listen
	for {

		// Reading what comes from the server
		n, err := r.Read(buf[:])

		if err != nil {
			return
		}

		println("Client got:", string(buf[0:n])) // Msg from server
	}
}

func main() {

	conn, err := net.Dial("unix", SOCK) // var SOCK = "/tmp/unix.sock"
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hello, I'm the client.\n")) // Sending msg to the server

	readerServer(conn) // Receiving from server msg

	if err != nil {
		log.Fatal("write error:", err)
	}

	time.Sleep(1e9)
}

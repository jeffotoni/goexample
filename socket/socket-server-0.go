/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"log"
	"net"
	"os"
)

var (
	SOCK = "/tmp/unix.sock"
)

func RedWriteSocket(conn net.Conn) {

	// Always Traveling and searching for messages
	for {

		buf := make([]byte, 512) // buffer
		nr, _ := conn.Read(buf)  // Reading msg coming from server
		data := buf[0:nr]        // data

		println("Server got:", string(data))

		///send client
		_, err = conn.Write([]byte("hello Client.. I'm the server, \n"))

		if err != nil {
			log.Fatal("Write: ", err)
		}
	}
}

func DeleteScock() {

	_, err := os.Stat(SOCK)
	if err == nil {
		if err = os.Remove(SOCK); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {

	DeleteScock()                      // remove socket
	l, err := net.Listen("unix", SOCK) // var SOCK = "/tmp/unix.sock"
	if err != nil {
		log.Fatal("listen error:", err)
	}

	for { // Connection-oriented

		conn, err := l.Accept() // Opening connection
		defer l.Close()

		if err != nil {
			log.Fatal("accept error:", err)
		}
		go RedWriteSocket(conn) // Receive msg from server and send msg to server
	}
}

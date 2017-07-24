package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	conn, err := listener.Accept()
	defer conn.Close()

	for {

		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message Received: \n", string(message))
	conn.Write([]byte("Hello:" + "\n")) // don't care about return value
	conn.Write([]byte(message + "\n"))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

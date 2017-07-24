/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var (
	PORT = ":7777"
)

func main() {

	tcp, err := net.Listen("tcp", PORT) // var PORT = :7777 or // ListenTCP("tcp", tcpAddr))
	fmt.Println("Open port: ", PORT)
	checkError(err)

	conn, _ := tcp.Accept()
	defer conn.Close()

	// send new string back to client
	conn.Write([]byte("Welcome at the call client!" + "\n"))

	// will listen for message to process ending in newline (\n)
	message, _ := bufio.NewReader(conn).ReadString('\n')

	// output message received
	fmt.Print("Client sent: \n", string(message))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

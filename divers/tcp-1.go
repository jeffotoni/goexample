/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size		   2017
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//netcat localhost 9001
//talk here
//
//or
//telnet localhost 9001
//talk here
//
//send msg
//echo "[MENSAGEM teste]" > /dev/tcp/localhost/9001
func main() {

	tcp, err := net.Listen("tcp", ":9001")

	if err != nil {
		log.Fatal(err)
	}

	conn, _ := tcp.Accept()
	defer conn.Close()

	// send new string back to client
	conn.Write([]byte("Welcome at the call!" + "\n"))

	// run loop forever (or until ctrl-c)
	for {

		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')

		if string(message) != "" {

			// output message received
			fmt.Print("Message Received: \n", string(message))

		} else {

			fmt.Println("close!!")
			break
		}
	}

	fmt.Println("down")
}

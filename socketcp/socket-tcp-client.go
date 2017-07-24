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
	"net"
	"os"
)

var (
	PORT = ":7777"
)

func main() {

	//Connect TCP
	conn, err := net.Dial("tcp", PORT) // var PORT = :7777
	checkError(err)

	defer conn.Close()

	//simple Read
	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)

	fmt.Println("server sent: ", string(buffer[0:n]))

	//simple write
	conn.Write([]byte("Hello, I'm the net client.\n"))

}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

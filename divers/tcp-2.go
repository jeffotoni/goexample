/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size		   2017
 */

package main

import (
	"fmt"
	"net"
)

func main() {

	//Connect TCP
	conn, err := net.Dial("tcp", "localhost:9001")
	if err != nil {

		fmt.Println("error: ", err)
	}

	defer conn.Close()

	//simple Read
	buffer := make([]byte, 1024)
	conn.Read(buffer)

	//simple write
	conn.Write([]byte("Hello from client"))
}

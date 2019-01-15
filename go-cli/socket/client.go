package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		//println("Client got:", string(buf[0:n]))
		println(string(buf[0:n]))
	}
}

func main() {

	c, err := net.Dial("unix", "/tmp/agent-123.sock")
	if err != nil {
		fmt.Println("not socket unix: ", err)
		return
	}
	defer c.Close()
	go reader(c)

	_, err = c.Write([]byte(`{"msg":"hello", "status":"success"}`))
	if err != nil {
		log.Println("write error waiting:", err)
		//time.Sleep(time.Second * 3)
		//break
	}
	time.Sleep(1e9)

	// for {
	// 	_, err := c.Write([]byte(`{"msg":"hello", "status":"success"}`))
	// 	if err != nil {
	// 		log.Println("write error waiting:", err)
	// 		//time.Sleep(time.Second * 3)
	// 		break
	// 	}
	// 	time.Sleep(1e9)
	// }
}

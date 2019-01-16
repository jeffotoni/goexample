// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
  "bufio"
  "fmt"
  "log"
  "net"
  "os"
)

const (
  HOST = "127.0.0.1"
  PORT = "8081"
)

func main() {

  // connect to this socket
  conn, _ := net.Dial("tcp", HOST+":"+PORT)

  if conn == nil {
    log.Println("server [" + HOST + ":" + PORT + "] not found")
    return
  }

  for {
    // read in input from stdin
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
    text, _ := reader.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, text+"\n")
    // listen for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message from server: " + message)
  }
}

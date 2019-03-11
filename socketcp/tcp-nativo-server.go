// Go in action
// @jeffotoni
// 2019-03-11

package main

import (
  "fmt"
  "net"
  "os"
)

//  8001 and 62554
const (
  CONN_HOST = "localhost"
  CONN_PORT = "22334"
  CONN_TYPE = "tcp"
)

func main() {
  // Listen for incoming connections.
  l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
  if err != nil {
    fmt.Println("Error listening:", err.Error())
    os.Exit(1)
  }
  // Close the listener when the application closes.
  defer l.Close()
  fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
  for {
    // Listen for an incoming connection.
    conn, err := l.Accept()
    if err != nil {
      fmt.Println("Error accepting: ", err.Error())
      os.Exit(1)
    }
    // Handle connections in a new goroutine.
    go handleRequest(conn)
  }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
  // Make a buffer to hold incoming data.
  buf := make([]byte, 1024)
  // Read the incoming connection into the buffer.
  reqLen, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading:", err.Error())
  }

  fmt.Println("reqLen: ", reqLen)
  fmt.Println("msg: ", string(buf))

  // Send a response back to person contacting us.
  conn.Write([]byte("ok"))
  // Close the connection when you're done with it.
  conn.Close()
}

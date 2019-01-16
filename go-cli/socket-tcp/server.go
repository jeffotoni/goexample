// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
  "bufio"
  "fmt"
  "net"
  "os"
  "strings"
)

// only needed below for sample processing

func main() {

  // listen on all interfaces
  ln, err := net.Listen("tcp4", ":8081")

  if err != nil {
    fmt.Println("Error listening:", err.Error())
    os.Exit(1)
  }
  // Close the listener when the application closes.
  //defer ln.Close()

  fmt.Println("Launching server :8081... ")

  // accept connection on port
  conn, err := ln.Accept()
  if err != nil {
    fmt.Println("Error accepting: ", err.Error())
    return
  }

  // run loop forever (or until ctrl-c)
  for {

    // will listen for message to process ending in newline (\n)
    message, _ := bufio.NewReader(conn).ReadString('\n')
    if len(string(message)) == 0 {
      continue
    }

    // output message received
    fmt.Print("Message Received:", string(message))
    // sample process for string received
    newmessage := strings.ToUpper(message)
    // send new string back to client
    conn.Write([]byte(newmessage + "\n"))
  }
}

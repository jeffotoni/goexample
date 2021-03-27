// Go in action
// @jeffotoni
// 2019-03-11

package main

import (
  "fmt"
  "log"
  "net"
  "os"
  "time"
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

    buf := make([]byte, 512)
    n, err := conn.Read(buf[0:])
    if err != nil {
      // handle error
      log.Println(err)
      time.Sleep(time.Second)
      continue
    }

    // message, _ := bufio.NewReader(conn).ReadString('\n')
    //println("size:", n, "msg:", string(buf))
    if n > 0 {
      jsonmsg := `{"versão": "1.1", "host": "exemplo.org", "short_message": "Uma mensagem curta", "nível": 5, "_some_info": "` + string(buf) + `"}`
      // Send a response back to person contacting us.
      conn.Write([]byte(jsonmsg))

    }
    // Handle connections in a new goroutine.
    //handleRequest(conn)

    // Close the connection when you're done with it.
    conn.Close()
    //time.Sleep(time.Second)
  }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
  // Make a buffer to hold incoming data.
  // buf := make([]byte, 1024)
  // // Read the incoming connection into the buffer.
  // reqLen, err := conn.Read(buf)
  // if err != nil {
  //   fmt.Println("Error reading:", err.Error())
  // }

  // fmt.Println("reqLen: ", reqLen)
  // fmt.Println("msg: ", string(buf))

  jsonmsg := `{"versão": "1.1", "host": "exemplo.org", "short_message": "Uma mensagem curta", "nível": 5, "_some_info": "foo"}`
  // Send a response back to person contacting us.
  conn.Write([]byte(jsonmsg))
  // Close the connection when you're done with it.
  conn.Close()
}

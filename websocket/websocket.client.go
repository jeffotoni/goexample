package main

import (
  "fmt"
  "log"
  "os"
  "time"

  "golang.org/x/net/websocket"
)

//type messageType struct{}

func main() {
  if len(os.Args) != 2 {
    log.Fatal("Uso: ./client [name]")
  }
  client := os.Args[1]

  origin := "http://localhost/"
  url := "ws://localhost:1234/chat"
  // i := 0
  for {
    //i++
    ws, err := websocket.Dial(url, "", origin)
    if err != nil {
      log.Fatal(err)
    }

    //istr := strconv.Itoa(i)
    if _, err := ws.Write([]byte(client)); err != nil {
      log.Fatal(err)
    }

    // for {
    var msg = make([]byte, 512)
    var n int
    if n, err = ws.Read(msg); err != nil {
      log.Fatal(err)
    }
    fmt.Printf("Received: %s.\n", msg[:n])
    time.Sleep(time.Millisecond * 300)
  }

  // // create connection
  // // schema can be ws:// or wss://
  // // host, port – WebSocket server
  // conn, err := websocket.Dial("ws://localhost:1234/chat", "", "http://localhost/")
  // if err != nil {
  //   // handle error
  //   log.Println(err)
  //   return
  // }
  // defer conn.Close()

  // // send message
  // if err = websocket.JSON.Send(conn, `{"msg":"jeffotoni", "code":"xxxxxxxxxxxxe39393"}`); err != nil {
  //   // handle error
  // }

  // // receive message
  // // messageType initializes some type of message
  // message := messageType{}
  // if err := websocket.JSON.Receive(conn, &message); err != nil {
  //   // handle error
  //   log.Println(err)
  // }
}

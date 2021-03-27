// Go in action
// @jeffotoni
// 2019-03-11

package main

import (
  "bufio"
  "log"
  "net"
  "os"
  "strings"
  "time"
)

func main() {

  // connect to this socket
  //conn, _ := net.Dial("tcp", "localhost:22334")
  //for {
  // read in input from stdin
  // reader := bufio.NewReader(os.Stdin)
  // fmt.Print("Text to send: ")
  // jsonmsg := `{"versão": "1.1", "host": "exemplo.org", "short_message": "Uma mensagem curta", "nível": 5, "_some_info": "foo"}`
  // text, _ := reader.ReadString('\n')
  // send to socket
  //fmt.Fprintf(conn, jsonmsg)
  // listen for reply

  if len(os.Args) != 2 {
    log.Fatal("Uso: ./client [name]")
  }

  str := os.Args[1]
  println("client v1.0.0")
  for {
    conn, _ := net.Dial("tcp", "localhost:22334")

    // var buf []byte
    // _, err := conn.Read(buf[0:])
    // if err != nil {
    //   // handle error
    //   log.Println(err)
    //   time.Sleep(time.Second)
    //   continue
    // }

    // send message
    _, err := conn.Write([]byte(`{"key":"x0000001_` + str + `"}`))
    if err != nil {
      // handle error
      log.Println(err)
      time.Sleep(time.Second)
      continue
    }

    message, _ := bufio.NewReader(conn).ReadString('\n')
    message = strings.Trim(message, " ")
    println(message)
    time.Sleep(time.Millisecond * 300)
  }

  // if strings.ToLower(message) == "ok" {
  //   fmt.Printf("\nSave")
  // } else {
  //   fmt.Printf("\nError server tcp")
  // }
  //}
}

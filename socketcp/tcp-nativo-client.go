// Go in action
// @jeffotoni
// 2019-03-11

package main

import (
  "bufio"
  "fmt"
  "net"
  "strings"
)

func main() {

  // connect to this socket
  conn, _ := net.Dial("tcp", "localhost:22334")
  //for {
  // read in input from stdin
  // reader := bufio.NewReader(os.Stdin)
  fmt.Print("Text to send: ")
  jsonmsg := `{"versão": "1.1", "host": "exemplo.org", "short_message": "Uma mensagem curta", "nível": 5, "_some_info": "foo"}`
  // text, _ := reader.ReadString('\n')
  // send to socket
  fmt.Fprintf(conn, jsonmsg)
  // listen for reply
  message, _ := bufio.NewReader(conn).ReadString('\n')
  message = strings.Trim(message, " ")
  if strings.ToLower(message) == "ok" {
    fmt.Printf("\nSave")
  } else {
    fmt.Printf("\nError server tcp")
  }
  //}
}

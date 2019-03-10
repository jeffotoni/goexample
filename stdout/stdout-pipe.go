// Go Api server
// @jeffotoni
// 2019-03-10

package main

import (
  "bytes"
  "fmt"
  "io"
  "os"
)

func main() {

  old := os.Stdout
  r, w, _ := os.Pipe()
  os.Stdout = w

  fmt.Println("Hello, playground") // this gets captured

  w.Close()
  //out, _ := ioutil.ReadAll(r)
  os.Stdout = old

  //fmt.Printf("Captured: %s", out) // prints: Captured: Hello, playground

  var buf bytes.Buffer
  io.Copy(&buf, r)
  fmt.Printf("Captured: %s", buf.String())
}

package main

import (
  "bytes"
  "fmt"
  "io"
  "os"
)

func main() {

  f := func() {
    fmt.Println("func here, stdout...")
  }

  out := captureStdout(f)
  fmt.Println(out)
}

// not thread safe
func captureStdout(f func()) string {
  old := os.Stdout
  r, w, _ := os.Pipe()
  os.Stdout = w

  f()

  w.Close()
  os.Stdout = old

  var buf bytes.Buffer
  io.Copy(&buf, r)
  return buf.String()
}

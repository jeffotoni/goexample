// Go in Action
// @jeffotoni
// 2019-03-25

package main

import (
  "fmt"
  "os"
)

func main() {
  fi, err := os.Stdin.Stat()
  if err != nil {
    panic(err)
  }
  if fi.Mode()&os.ModeNamedPipe == 0 {
    fmt.Println("no pipe :(")
  } else {
    fmt.Println("stdin is empty")
  }
}

package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

func main() {

  rescueStdout := os.Stdout
  r, w, _ := os.Pipe()
  os.Stdout = w

  fmt.Println("Hello, playground") // this gets captured
  fmt.Println("I'm testing")       // this gets captured
  fmt.Println("Let's Go..")        // this gets captured

  w.Close()
  out, _ := ioutil.ReadAll(r)
  os.Stdout = rescueStdout

  //fmt.Sprintf("%s", out)
  fmt.Printf("Captured: %s", out) // prints: Captured: Hello, playground
}

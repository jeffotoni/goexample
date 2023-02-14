// Go Api server
// @jeffotoni
// 2019-03-10

package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
)

func print() {
    fmt.Println("output")
}

func main() {
    // setting stdout to a file
    fname := filepath.Join(os.TempDir(), "stdout")
    fmt.Println("stdout is now set to", fname)
    old := os.Stdout            // keep backup of the real stdout
    temp, _ := os.Create(fname) // create temp file
    os.Stdout = temp

    // capture
    print()

    // back to normal state
    temp.Close()
    os.Stdout = old // restoring the real stdout

    // reading our temp stdout
    fmt.Println("previous output:")
    out, _ := ioutil.ReadFile(fname)
    fmt.Print(string(out)
}

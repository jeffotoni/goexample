// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    p := make([]byte, 2048)
    conn, err := net.Dial("udp", "127.0.0.1:1234")
    if err != nil {
        fmt.Printf("Some error %v", err)
        return
    }
    fmt.Fprintf(conn, `{"cpu":"34.44","memory":"88.99"}`)
    _, err = bufio.NewReader(conn).Read(p)
    if err == nil {
        fmt.Printf("%s\n", p)
    } else {
        fmt.Printf("Some error %v\n", err)
    }
    conn.Close()
}

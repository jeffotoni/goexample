// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
    "fmt"
    "net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
    _, err := conn.WriteToUDP([]byte(`{"status":"success", "msg":"sending message via upd"}`), addr)
    if err != nil {
        fmt.Printf("Couldn't send response %v", err)
    }
}

func main() {
    p := make([]byte, 2048)
    addr := net.UDPAddr{
        Port: 1234,
        IP:   net.ParseIP("127.0.0.1"),
    }
    ser, err := net.ListenUDP("udp", &addr)
    if err != nil {
        fmt.Printf("Some error %v\n", err)
        return
    }

    fmt.Println("Listen Server Udp *:1234")
    for {
        _, remoteaddr, err := ser.ReadFromUDP(p)
        fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
        if err != nil {
            fmt.Printf("Some error  %v", err)
            continue
        }
        go sendResponse(ser, remoteaddr)
    }
}

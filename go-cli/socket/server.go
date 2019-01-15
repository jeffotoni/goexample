package main

import (
    "log"
    "net"
    "os"
)

func echoServer(c net.Conn) {
    for {
        buf := make([]byte, 512)
        nr, err := c.Read(buf)
        if err != nil {
            return
        }

        data := buf[0:nr]
        println("Server got:", string(data))
        _, err = c.Write(data)
        if err != nil {
            log.Fatal("Write: ", err)
        }
    }
}

func main() {

    socket := "/tmp/agent-123.sock"
    // rm file
    // /tmp/agent-123.sock
    os.Remove(socket)
    l, err := net.Listen("unix", socket)
    if err != nil {
        log.Fatal("listen error:", err)
    }

    for {
        fd, err := l.Accept()
        if err != nil {
            log.Fatal("accept error:", err)
        }

        go echoServer(fd)
    }
}

package main

import (
    "bytes"

    "fmt"
    "io"
    "log"

    "net/http"
    "os"

    "github.com/golang/protobuf/proto"
    protoc "github.com/jeffotoni/go.protobuffer.customer"
)

func main() {
    body := &protoc.Customer{
        Id:   12304,
        Name: "Carlos",
    }

    data, err := proto.Marshal(body)
    if err != nil {
        log.Println("proto.Marshal", err)
        return
    }

    req, err := http.NewRequest("POST", "http://localhost:8080/customer/proto", bytes.NewBuffer(data))
    req.Header.Set("Content-Type", "application/proto")

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        log.Println("Error Do:", err)
        return
    }

    defer res.Body.Close()
    fmt.Println("response Status:", res.Status)
    io.Copy(os.Stdout, res.Body)
}

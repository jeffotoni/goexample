package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

type Todo struct {
    Name   string `json:"name"`
    Cpf    string `json:"cpf"`
    Status string `json:"status"`
}

func main() {
    deleteHttp()
}

func deleteHttp() {
    fmt.Println("Http Delete...")
    todo := Todo{"jeffotoni", "123.393.334-39", "ativo"}
    jsonReq, err := json.Marshal(todo)
    if err != nil {
        log.Println("error marshal:", err.Error())
        return
    }

    // strings.NewReader(reqBody)
    req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/api/v1/user", bytes.NewBuffer(jsonReq))
    if err != nil {
        log.Println("error NewRequest:", err.Error())
        return
    }

    client := &http.Client{}
    req.Header.Add("Content-Type", "application/json")
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalln("error:", err)
    }

    defer resp.Body.Close()
    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("error ReadAll:", err.Error())
        return
    }

    bodyString := string(bodyBytes)
    fmt.Println("body:", bodyString)
}

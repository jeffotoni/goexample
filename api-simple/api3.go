// Go in action
// @jeffotoni
// 2019-04-05

/// curl localhost:8080/api/hello
//////////////////////////////////////

package main

import (
    //"fmt"
    "log"
    "net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {

    switch r.Method {
    case http.MethodPost:
        log.Println("Method: ", http.MethodPost)
        break
    case http.MethodGet:
        log.Println("Method: ", http.MethodGet)
        break
    case http.MethodPut:
        log.Println("Method: ", http.MethodPut)
        break
    case http.MethodDelete:
        log.Println("Method: ", http.MethodDelete)
        break
    default:
        log.Println("Method não permitido: ", r.Method)
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello, welcome to the world, Go!"))
}

func main() {

    mux := http.NewServeMux()

    //mux.Handle("/api/ping", http.HandlerFunc(Hello))
    mux.Handle("/api/hello", http.HandlerFunc(Hello))
    // mux.Handle("/api/products", http.HandlerFunc(Hello))
    // mux.Handle("/api/products/{id}", http.HandlerFunc(Hello))

    server :=
        &http.Server{
            Addr:    ":8080",
            Handler: mux,
        }

    log.Println("Server Run port: 8080")
    if err := server.ListenAndServe(); err != nil {
        log.Printf("Eror while serving metrics: %s", err)
    }
}

// Go in action
// @jeffotoni
// 2019-04-05

/// curl localhost:8080/api/hello
//////////////////////////////////////

package main

import (
    "fmt"
    "log"
    "net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
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
            Addr:    ":8888",
            Handler: mux,
        }

    fmt.Println("Run Server...")
    if err := server.ListenAndServe(); err != nil {
        log.Printf("Eror while serving metrics: %s", err)
    }
}

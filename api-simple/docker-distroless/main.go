// Go in action
// @jeffotoni
// 2019-04-05

/// docker build -t example-distroless -f Dockerfile .
/// docker run --rm -it example-distroless
/// curl localhost:8080/api/hello

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
    mux.Handle("/api/hello", http.HandlerFunc(Hello))

    server :=
        &http.Server{
            Addr:    ":8080",
            Handler: mux,
        }

    fmt.Printf("Server Run port: 8080\n")
    if err := server.ListenAndServe(); err != nil {
        log.Printf("Eror while serving metrics: %s", err)
    }
}

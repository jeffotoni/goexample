// Go in action
// @jeffotoni

/// curl localhost:8080/user
//////////////////////////////////////

package main

import (
    "github.com/rs/cors"
    "log"
    "net/http"
)

func User(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"user":"jeffotoni","email":"jeffotoni@gmail.com"}`))
}

func main() {

    mux := http.NewServeMux()
    mux.Handle("/user", http.HandlerFunc(User))

    corsmux := cors.Default().Handler(mux)

    server :=
        &http.Server{
            Addr:    ":8080",
            Handler: corsmux,
        }
    log.Println("Server Run port: 8080")
    if err := server.ListenAndServe(); err != nil {
        log.Printf("Error server: %s", err)
    }
}

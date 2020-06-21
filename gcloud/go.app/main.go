package main

import (
        "log"
        "net/http"
        "os"
)

var (
        port string
)

func init() {
        port = os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }
}

func main() {

        mux := http.NewServeMux()
        mux.HandleFunc("/ping", ping)
        log.Printf("Server listening on port %s", port)
        log.Fatal(http.ListenAndServe(":"+port, mux))
}

func ping(w http.ResponseWriter, r *http.Request) {
        log.Printf("Serving request: %s", r.URL.Path)
        host, _ := os.Hostname()
        jres := `{"resp":"pong", "version":"1.0.0", "hostname":"` + host + `"}`
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(jres))
}

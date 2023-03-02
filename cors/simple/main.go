package main

import (
    "io"
    "net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    b, err := io.ReadAll(r.Body)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte(`{"msg":"error"}`))
        return
    }
    w.WriteHeader(200)
    w.Write(b)
}

func OtherHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Outro endpoint!"))
}

func Cors(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Vary", "Origin")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
        next.ServeHTTP(w, r)
    })
}

func main() {
    mux := http.NewServeMux()
    mux.Handle("/v1/user", &MyHandler{})
    mux.HandleFunc("/outro", OtherHandler)

    newmux := Cors(mux)
    http.ListenAndServe(":8080", newmux)
}

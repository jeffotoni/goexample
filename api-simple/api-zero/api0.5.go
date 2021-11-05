package main

import (
	"log"
	"net/http"
)

func main() {
	println("Run Server Port:5010")
	http.HandleFunc("/api/v1/ping", Ping)
	log.Println(http.ListenAndServe("0.0.0.0:5010", nil))
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg":"ok}`))
}

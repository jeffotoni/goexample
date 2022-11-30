package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", Hello)
	log.Println("Run Server port:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"msg":"success"}`))
	return
}

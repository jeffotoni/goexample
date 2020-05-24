package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

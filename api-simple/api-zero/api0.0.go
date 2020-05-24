package main

import (
	"log"
	"net/http"
)

func main() {

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello Gophers..."))
		})))
}

package main

import (
	"log"
	"net/http"
)

func main() {
	println("\nServer run 8080\n")
	err := http.ListenAndServe("0.0.0.0:8080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("est Request One!!"))
		}))
	log.Fatal(err)
}

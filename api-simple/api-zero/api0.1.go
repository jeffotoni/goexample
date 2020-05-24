package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("\nServer run 8080\n")
	err := http.ListenAndServe(":8080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("DevopsBH for Golang simple" + r.URL.Path))
		}))
	log.Fatal(err)
}

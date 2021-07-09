package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/ping1",
		func(w http.ResponseWriter, r *http.Request) {
			println("ok")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong 1"))
		})
	mux.HandleFunc("/ping2",
		func(w http.ResponseWriter, r *http.Request) {
			println("ok")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong 2"))
		})

	mux.HandleFunc("/ping3",
		func(w http.ResponseWriter, r *http.Request) {
			println("ok")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong 3"))
		})

	mux.HandleFunc("/ping4",
		func(w http.ResponseWriter, r *http.Request) {
			println("ok")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong 4"))
		})

	mux.HandleFunc("/ping5",
		func(w http.ResponseWriter, r *http.Request) {
			println("ok")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong 5"))
		})

	mux.HandleFunc("/ping6",
		func(w http.ResponseWriter, r *http.Request) {
			println("ok")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong 05"))
		})

	mux.HandleFunc("/ping7",
		func(w http.ResponseWriter, r *http.Request) {
			println("ok")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong 07"))
		})
	mux.HandleFunc("/ping8",
		func(w http.ResponseWriter, r *http.Request) {
			println("ok")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong 08"))
		})

	mux.HandleFunc("/ping9",
		func(w http.ResponseWriter, r *http.Request) {
			println("ok")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong 09"))
		})

	mux.HandleFunc("/ping10",
		func(w http.ResponseWriter, r *http.Request) {
			println("ok")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong 10"))
		})

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	log.Print("Run Server:8080")
	log.Fatal(server.ListenAndServe())
}

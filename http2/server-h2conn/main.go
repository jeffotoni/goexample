package main

import (
	"io"
	"log"
	"net/http"

	"github.com/posener/h2conn"
)

// 	openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt
func main() {
	srv := &http.Server{Addr: ":8000", Handler: http.HandlerFunc(echo)}
	log.Print("Echo on https://localhost:8000")
	log.Print("Run curl -k -i --http2 https://localhost:8000 -d test")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

func echo(w http.ResponseWriter, r *http.Request) {
	// Accept returns a connection to the client  that can be used:
	//   1. Write - send data to the client
	//   2. Read - receive data from the client
	conn, err := h2conn.Accept(w, r)
	if err != nil {
		log.Printf("Failed creating connection from %s: %s", r.RemoteAddr, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Send back to the client everything that we receive
	io.Copy(conn, conn)
}

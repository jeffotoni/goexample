package main

import (
	// "fmt"
	// "io"
	"log"
	"net/http"
)

// gerando server key
// openssl ecparam -genkey -name secp384r1 -out server.key

// gerando server crt
// openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func main() {

	http.HandleFunc("/hello", HelloServer)

	err := http.ListenAndServeTLS("0.0.0.0:443", "server.crt", "server.key", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size		   2017
 */

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"

	"github.com/gorilla/mux"
)

const (
	SOCK = "/tmp/tcp.sock"
)

func homeView(w http.ResponseWriter, r *http.Request) {

	headers := w.Header()

	headers.Add("Content-Type", "text/html")

	io.WriteString(w, "<html><head></head><body><p>It works!</p></body></html>")
}

// type IP []byte
// which are TCP sockets and UDP sockets.
func main() {

	if err := os.Remove(SOCK); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", homeView)

	fmt.Println("Listen Unix: ", SOCK)

	listener, err := net.Listen("unix", SOCK)

	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	err = fcgi.Serve(listener, r)
	if err != nil {

		fmt.Println("error socker: ", err)
	}
}

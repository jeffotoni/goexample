/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size		   2017
 */

package main

import (
	"bufio"
	"fmt"
	"html"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func Xhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// My func handler
func MyFuncHandler(w http.ResponseWriter, r *http.Request) {

	if html.EscapeString(r.URL.Path) != "/postest" {

		fmt.Fprintf(w, "500", "Nao atorizado!\n")
	} else {

		// Sends to the client
		fmt.Fprintf(w, "Sua url: %q\n", html.EscapeString(r.URL.Path))
	}

}

const (
	SOCK = "/tmp/go.sock"
)

type Server struct {
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	body := "Hello World\n"
	w.Header().Set("Server", "gophr")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprint(len(body)))
	fmt.Fprint(w, body)
}

func main() {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, syscall.SIGTERM)

	server := Server{}
	//http.Handle("/", server)
	//http.HandleFunc("/", MyFuncHandler)
	//http.HandleFunc("/postest", MyFuncHandler)
	//http.ListenAndServe(":8080", nil)

	go func() {

		//http.Handle("/", server)
		http.HandleFunc("/postest", MyFuncHandler)
		if err := http.ListenAndServe(":8088", nil); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		tcp, err := net.Listen("tcp", ":9001")

		///
		conn, _ := tcp.Accept()

		if err != nil {

			log.Fatal(err)
		}

		// run loop forever (or until ctrl-c)
		for {
			// will listen for message to process ending in newline (\n)
			message, _ := bufio.NewReader(conn).ReadString('\n')
			// output message received
			fmt.Print("Message Received:", string(message))
			// sample process for string received
			newmessage := strings.ToUpper(message)
			// send new string back to client
			conn.Write([]byte(newmessage + "\n"))
		}

		fcgi.Serve(tcp, server)
	}()

	<-sigchan

	if err := os.Remove(SOCK); err != nil {
		log.Fatal(err)
	}

}

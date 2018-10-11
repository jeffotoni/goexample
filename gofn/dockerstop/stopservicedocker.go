package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	shutdown := make(chan int)

	//create a notification channel to shutdown
	sigChan := make(chan os.Signal, 1)

	server := &http.Server{

		Addr: ":8080",
		// timeout para o nosso http
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	//start the http server
	http.HandleFunc("/", hello)

	// derrubando servico
	http.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Shutting down now...")
		server.Close()
	})

	//server := manners.NewWithServer(&http.Server{Addr: ":8080", Handler: nil})

	fmt.Println("start service 8080")

	go func() {
		server.ListenAndServe()
		shutdown <- 1
	}()

	//register for interupt (Ctrl+C) and SIGTERM (docker)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("Shutting down...")
		server.Close()
	}()

	<-shutdown
}

func hello(w http.ResponseWriter, r *http.Request) {
	//pretend to do some work
	time.Sleep(300 * time.Millisecond)
	io.WriteString(w, "Hello world!")
}

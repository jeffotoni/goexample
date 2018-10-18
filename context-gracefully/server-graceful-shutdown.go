package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	service = "fooapi"
)

var (
	RESTART_VALOR   = 10
	port            = flag.Int("port", 3333, "http port to listen on")
	shutdownTimeout = flag.Duration("shutdown-timeout", 10*time.Second,
		"shutdown timeout (5s,5m,5h) before connections are cancelled")
)

func init() {

	// alterar valor
	// em memoria
	RESTART_VALOR = 200
}

func main() {

	flag.Parse()

	mux := http.NewServeMux()
	//mux.HandleFunc("/v1/api/hello", HandleIndex)

	mux.HandleFunc("/v1/api/hello", func(w http.ResponseWriter, r *http.Request) {
		//io.WriteString(w, "ok value="+fmt.Sprintf("%d", RESTART_VALOR))
		//w.WriteHeader(200)
		go HandleIndex(w, r)
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: mux,
	}

	sigc := make(chan os.Signal, 1)
	sigc2 := make(chan os.Signal, 1)

	signal.Notify(sigc, os.Interrupt)
	signal.Notify(sigc, syscall.SIGTERM)
	signal.Notify(sigc, syscall.SIGHUP)

	//signal.Notify(sigc, os.Interrupt)
	// signal.Notify(sigc,
	// 	os.Interrupt,
	// 	syscall.SIGHUP,
	// 	syscall.SIGINT,
	// 	syscall.SIGTERM,
	// 	syscall.SIGQUIT)

	// log.Printf("%s listening on 0.0.0.0:%d with %v timeout", service, *port, *shutdownTimeout)
	// if err := srv.ListenAndServe(); err != nil {
	// 	if err != http.ErrServerClosed {
	// 		log.Fatal(err)
	// 	}
	// }

	go func() {
		log.Printf("%s listening on 0.0.0.0:%d with %v timeout and restart value = %d", service, *port, *shutdownTimeout, RESTART_VALOR)
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal("log err: ", err)
			}
		}
	}()

	go func() {
		sig := <-sigc
		ss := fmt.Sprintf("%s here: ", sig)
		v := strings.Split(ss, " ")
		//fmt.Println("Sutting down gracefully:: ", sig)
		log.Println("gracefully 0::", v[0])

		if v[0] == "hangup" {

			RESTART_VALOR = 0
			log.Println("restrt processo...")

		} else {

			// clean up here
			fmt.Println("stop")
			// log.Printf("%s shutting down ...\n")
			ctx, cancel := context.WithTimeout(context.Background(), *shutdownTimeout)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("context: ", err)
			}

			os.Exit(0)
		}
	}()

	<-sigc2

	log.Printf("%s down\n", service)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, 世界"))
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {

	RESTART_VALOR = RESTART_VALOR + 100
	w.WriteHeader(200)
	io.WriteString(w, "ok value="+fmt.Sprintf("%d", RESTART_VALOR))
	//w.Write([]byte("ok"))
}

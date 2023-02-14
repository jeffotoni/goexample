package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	//work here

	go gracefulShutdown()
	forever := make(chan int)
	<-forever
}

func gracefulShutdown() {

	s := make(chan os.Signal, 1)

	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	signal.Notify(s, syscall.SIGHUP)

	go func() {

		sig := <-s
		ss := fmt.Sprintf("%s", sig)
		v := strings.Split(ss, " ")
		//fmt.Println("Sutting down gracefully:: ", sig)
		fmt.Println("gracefully 0::", v[0])

		if v[0] == "hangup" {

			fmt.Println("restrt")

		} else {

			// clean up here
			fmt.Println("stop")
			os.Exit(0)
		}

		//fmt.Println("gracefully 1:: ", v[1])

	}()
}

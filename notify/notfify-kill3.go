package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Program that will listen to the SIGINT and SIGTERM
// SIGINT will listen to CTRL-C.
// SIGTERM will be caught if kill command executed.
//
// See:
// - https://en.wikipedia.org/wiki/Unix_signal
// - https://www.quora.com/What-is-the-difference-between-the-SIGINT-and-SIGTERM-signals-in-Linux
// - http://programmergamer.blogspot.co.id/2013/05/clarification-on-sigint-sigterm-sigkill.html
func main() {
	errc := make(chan error)

	go func() {
		log.Println("Listening signals...")
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("Signal %v", <-c)
	}()

	log.Println("Exit:", <-errc)
}

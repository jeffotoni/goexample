package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time" // or "runtime"
)

func close() {
	// voltando o cursor
	fmt.Print("\033[?25h")
	fmt.Println("fechando programa")
}

func main() {

	go func() {

		c := make(chan os.Signal, 2)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		close()
		os.Exit(0)
	}()

	go func() {
		// removendo cursor
		fmt.Print("\033[?25l")
		timer := time.Tick(time.Duration(50) * time.Millisecond)

		s := []rune(`|/~\`)
		i := 0

		for {

			<-timer
			fmt.Print("\r")
			fmt.Print("\033[0;33m" + string(s[i]) + "\033[0m")

			i++
			if i == len(s) {
				i = 0
			}

			//time.Sleep(1 * time.Second) // or runtime.Gosched() or similar per @misterbee
		}
	}()

	// executando
	// programa
	// em tempos em tempos
	for {

		<-time.After(time.Duration(3 * time.Second))
		fmt.Print("   => executa..")
		fmt.Print("\r")
	}
}

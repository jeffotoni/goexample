package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	go func() {
		lambda()
	}()

	time.Sleep(20 * time.Second)
	fmt.Println("Shutting down...")
	os.Exit(0)
}

func lambda() {

	fmt.Println("Executando Lambda")

	for {

		time.Sleep(1800 * time.Millisecond)
	}
}

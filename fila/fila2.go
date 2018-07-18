package main

import (
	"fmt"
	"time"
)

var (
	token chan bool
)

func main() {
	tokens := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		tokens <- true
	}

	for i := 0; i < 100; i++ {
		go func(i int) {
			<-tokens
			fmt.Printf("client %d is running\n", i)
			time.Sleep(time.Second)
			tokens <- true
		}(i)
	}

	// unorthodox waitgroup
	time.Sleep(time.Minute)
}

package main

import (
	"fmt"
	"time"
)

func send(out, finish chan bool) {
	for {
		out <- true
		fmt.Println("Write")
		time.Sleep(time.Millisecond * 100)
	}
}

func write(in chan bool) {
	for range in {
		fmt.Println("Read")
	}
}

func main() {
	chanFoo := make(chan bool)
	done := make(chan bool)

	go send(chanFoo, done)
	go write(chanFoo)

	<-done
}

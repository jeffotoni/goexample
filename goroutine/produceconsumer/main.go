package main

import (
	"fmt"
	"time"
)

var done = make(chan bool)
var msgs = make(chan int)

func produce() {
	for i := 0; i < 4; i++ {
		fmt.Println("sending")
		msgs <- i
		fmt.Println("sent")
	}
	fmt.Println("Before closing channel")
	close(msgs)
	fmt.Println("Before passing true to done")
	done <- true
}

func consume() {
	for msg := range msgs {
		fmt.Println("Consumer: ", msg)
		time.Sleep(100 * time.Millisecond)

	}
}

func main() {
	go produce()
	go consume()
	<-done
	fmt.Println("After calling DONE")
}

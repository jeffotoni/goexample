package main

import (
	"fmt"
	"syscall"
)

func block(c chan bool) {

	fmt.Println("block() enter")
	buf := make([]byte, 1024)
	_, _ = syscall.Read(0, buf) // block on doing an unbuffered read on STDIN
	fmt.Println("block() exit")
	c <- true // main() we're done
}

func main() {

	c := make(chan bool)
	for i := 0; i < 1000; i++ {
		go block(c)
	}
	for i := 0; i < 1000; i++ {
		_ = <-c
	}
}

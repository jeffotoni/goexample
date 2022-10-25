package main

import (
	"fmt"
)

func hello(done chan int) {
	fmt.Println("hello func")
	done <- 5
}
func main() {
	done := make(chan int, 1)
	done <- 3
	fmt.Println("Hello, playground")
	//go hello(done)
	fmt.Println("Hello done", <-done)
}

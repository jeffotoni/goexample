package main

import (
	"fmt"
	"os"
	"strconv"
)

var length int

func Fibonacci(count int) {
	ch := make(chan int)
	shouldQuit := make(chan bool)

	go func() {

		for i := 0; i < count; i++ {
			n := <-ch
			fmt.Printf("%d, \n", n)
		}

		shouldQuit <- true

	}()

	x, y := 0, 1

	for {

		select {
		case ch <- x:
			x, y = y, x+y
		case isDone := <-shouldQuit:
			fmt.Println("\nisDoen?", isDone)
			return
		}
	}

}

func GetLength() {

	n := 12

	if len(os.Args) == 2 {
		n, _ = strconv.Atoi(os.Args[1])
	}

	length = n

}

func main() {

	GetLength()
	Fibonacci(length)

}

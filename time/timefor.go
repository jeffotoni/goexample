package main

import (
	"fmt"
	"time"
)

var c chan int

func handle(int) { println("handle.. ") }

func main() {

	var i int

	//go func() {
	for {
		//select {
		//case m := <-c:
		//handle(m)
		//case
		<-time.After(4 * time.Second)
		i++
		fmt.Println("timed out ", i)
		println("continua...")
	}

	//}()
}

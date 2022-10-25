package main

import "fmt"

func main() {

	c := make(chan int)

	// resolve com buffer
	// c := make(chan int, 1)

	// resolve com goroutine
	//go func() {
	c <- 1
	//}()

	fmt.Println(<-c)

}

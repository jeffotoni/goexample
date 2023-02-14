package main

import (
	"fmt"
)

var j = 10

func main() {
	// The code is blocked until something gets pushed into the returned channel
	// As opposed to the previous method, we block in the main function, instead
	// of the function itself
	i := <-getNumberChan()
	fmt.Println(i)
}

// return an integer channel instead of an integer
func getNumberChan() <-chan int {

	// create the channel
	c := make(chan int)

	go func() {

		fmt.Println(j)

		j++

		// push the result into the channel
		c <- j
	}()

	// se nao fechar
	// ele ocorre race warning
	<-c

	go func() {

		fmt.Println(j)

		j++

		// push the result into the channel
		c <- j
	}()

	// immediately return the channel
	return c
}

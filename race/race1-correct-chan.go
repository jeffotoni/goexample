package main

import (
	"fmt"
)

var j = 10

func main() {
	fmt.Println(getNumber())
}

func getNumber() int {
	var i int

	// Create a channel to push an empty struct to once we're done
	done := make(chan struct{})

	go func() {
		i = 5

		fmt.Println(j)

		j++
		// Push an empty struct once we're done
		done <- struct{}{}

	}()
	// This statement blocks until something gets pushed into the `done` channel
	<-done

	go func() {
		i++

		fmt.Println(j)

		j++
		// Push an empty struct once we're done
		done <- struct{}{}

	}()
	// This statement blocks until something gets pushed into the `done` channel
	<-done

	return i
}

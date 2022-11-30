package main

import (
	"fmt"
	"time"
)

func main() {

	// goroutine
	go func() {
		for {
			fmt.Println("Infinite Loop 1")
			time.Sleep(time.Second)
		}
	}()

	// no goroutine
	for true {
		fmt.Println("Infinite Loop 2")
		time.Sleep(time.Second)
	}
}

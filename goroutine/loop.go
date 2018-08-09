package main

import (
	"fmt"
	"time"
)

func main() {

	// while true loop
	for {
		fmt.Println("Infinite Loop 1")
		time.Sleep(time.Second)
	}

	// Alternative Version
	for true {
		fmt.Println("Infinite Loop 2")
		time.Sleep(time.Second)
	}
}

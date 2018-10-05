package main

import (
	"fmt"
)

func main() {

	ch := make(chan string, 2)
	ch <- "DevOpsBh"
	ch <- "2018 Bh"
	ch <- "Deadlock"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

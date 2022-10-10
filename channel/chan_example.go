package main

import "log"

func printResult() {

	x := generateChan(10)
	y := generateChan(15)

	result := increment(x, y)

	for r := range result {
		log.Printf("\ndata: %v", r)
	}
}

func increment(x ...chan int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range x {
			val := <-n
			val++
			out <- val
		}
		close(out)
	}()
	return out
}

func generateChan(x int) chan int {
	out := make(chan int)

	go func() {
		out <- x
		close(out)
	}()

	return out
}

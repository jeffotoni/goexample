package main

import "fmt"

func main() {

	c := fatorial(4)
	for v := range c {
		fmt.Println(v)
	}
}

func fatorial(n int) chan int {

	out := make(chan int)
	// basic
	go func() {

		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}

		out <- total
		close(out)
	}()

	return out
}

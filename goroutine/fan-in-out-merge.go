package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen(2, 3)

	// Distribui sq Work trabalhar em duas goroutine que tanto ler em
	// FAN-OUT
	c1 := sq(in)
	c2 := sq(in)

	// consume e da merge em saida de c1 e c2
	// // FAN-IN
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) <-chan int {

	fmt.Println("Type of nums %T\n", nums)
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {

	out := make(chan int)
	go func() {
		for n := range in {

			out <- n * n
		}
		close(out)
	}()

	return out
}

func merge(cs ...<-chan int) chan int {

	fmt.Printf("Type of cs %T\n", cs)

	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))

	for _, c := range cs {

		go output(c)

		// go func(ch chan int) {
		// 	for n := range ch {
		// 		out <- n
		// 	}
		// 	wg.Done()
		// }(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

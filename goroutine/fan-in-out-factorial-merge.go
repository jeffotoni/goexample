package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen()

	// FAN OUT
	// Multiplas funcoes
	// lendo um canal
	f1 := fatorial(in)
	f2 := fatorial(in)
	f3 := fatorial(in)
	f4 := fatorial(in)
	f5 := fatorial(in)
	f6 := fatorial(in)
	f7 := fatorial(in)
	f8 := fatorial(in)
	f9 := fatorial(in)

	//FAN IN
	// MULTIPLOS CANAIS CONVERTENDO EM UM UNICO CANAL
	for n := range merge(f1, f2, f3, f4, f5, f6, f7, f8, f9) {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 1; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()

	return out
}

func fatorial(in <-chan int) <-chan string {
	out := make(chan string)
	go func() {
		for n := range in {
			out <- fmt.Sprintf("fatorial %d = %d", n, fact(n))
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan string) chan string {

	fmt.Printf("Type of cs %T\n", cs)

	var wg sync.WaitGroup
	out := make(chan string)

	output := func(c <-chan string) {
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

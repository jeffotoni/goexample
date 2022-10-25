package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	N int = 40
)

func FactorialClosure() func(n uint64) uint64 {
	var a, b uint64 = 1, 1

	return func(n uint64) uint64 {
		if n > 1 {
			a, b = uint64(b), uint64(n)*uint64(a)
		} else {
			return 1
		}
		return b
	}
}

func main() {

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(100)
	}

	fmt.Println("FACTORIAL GOROUTINE")

	time.Sleep(time.Millisecond * 100)

	fact := FactorialClosure()

	start := time.Now()

	for i := uint64(0); i < uint64(N); i++ {
		go func(v uint64) {
			fmt.Printf("Factorial for %d is : %d \n", uint64(v), fact(uint64(v)))
		}(i)
	}

	end := time.Now()
	fmt.Printf("Tempo gasto %s\n", end.Sub(start))
}

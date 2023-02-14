package main

import (
	"fmt"
)

func fib(n int64) int64 {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-2) + fib(n-1)
}

func memoize(targetFunc func(int64) int64, cache map[int64]int64) func(int64) int64 {

	middlelayer := func(n int64) int64 {
		if cache[n] != 0 {
			return cache[n]
		}
		return targetFunc(n)
	}
	return func(n int64) int64 {
		for cache[n] == 0 {
			cache[n] = middlelayer(n-1) + middlelayer(n-2)
		}
		return cache[n]
	}
}
func main() {
	cache := make(map[int64]int64)
	cache[1] = 1
	cache[2] = 1
	newfib := memoize(fib, cache)
	for i := int64(1); i < 100; i++ {
		fmt.Printf("%d \n", newfib(i))
	}
}

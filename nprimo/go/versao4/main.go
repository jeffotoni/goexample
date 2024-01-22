package main

import (
	"fmt"
	"time"
)

func sieveOfEratosthenes(n int) []int {
	primes := make([]bool, n+1)
	for i := range primes {
		primes[i] = true
	}
	primes[0], primes[1] = false, false

	for p := 2; p*p <= n; p++ {
		if primes[p] {
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	var result []int
	for p, prime := range primes {
		if prime {
			result = append(result, p)
		}
	}

	return result
}

func main() {
	n := 1000000
	start := time.Now()
	primes := sieveOfEratosthenes(n)
	duration := time.Since(start)

	fmt.Println("Prime numbers up to", n, "are:")
	for _, prime := range primes {
		fmt.Println(prime)
	}
	fmt.Println("Time elapsed:", duration)
}

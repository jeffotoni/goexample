package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	primes := make([]int, 0)
	for i := 2; i <= 1000000; i++ {
		if i%2 == 0 {
			if i == 2 {
				 primes = append(primes, i)
			} else {
				continue
			}
		}

		isPrime := true
		for j := 3; j*j <= i; j += 2 {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
			fmt.Println(i)
		}
	}

	for _, prime := range primes {
		fmt.Println(prime)
	}

	end := time.Now()
	finish := end.Sub(start)
	fmt.Println("time:", finish)
}

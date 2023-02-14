package main

import (
	"fmt"
	"time"
)

const LIMIT = 41

var vetfact [LIMIT]uint64

func main() {

	fmt.Println("MEMOIZATION")

	start := time.Now()
	var result uint64 = 0

	for i := uint64(0); i < LIMIT; i++ {

		result = FactorialMemoization(uint64(i))
		fmt.Printf("O factorial %d é %d\n", uint64(i), uint64(result))
	}

	end := time.Now()
	fmt.Printf("Tempo gasto %s\n", end.Sub(start))
	fmt.Println("FIM MEMOIZATION")
}

func FactorialMemoization(n uint64) (res uint64) {

	if vetfact[n] != 0 {
		res = vetfact[n]
		return res
	}

	if n > 0 {

		res = n * FactorialMemoization(n-1)

		vetfact[n] = res

		return res
	}

	return 1
}

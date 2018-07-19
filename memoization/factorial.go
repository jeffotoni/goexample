package main

import (
	"fmt"
	"time"
)

const LIM = 41

var facts [LIM]uint64

func main() {

	fmt.Println("==================FACTORIAL==================")
	start := time.Now()
	for i := uint64(0); i < LIM; i++ {
		fmt.Printf("Factorial for %d is : %d \n", i, Factorial(uint64(i)))
	}
	end := time.Now()
	fmt.Printf("Calculation finished in %s \n", end.Sub(start)) //Calculation finished in 2.0002ms

	fmt.Println("==================FACTORIAL CLOSURE==================")
	start = time.Now()
	fact := FactorialClosure()
	for i := uint64(0); i < LIM; i++ {
		fmt.Printf("Factorial closure for %d is : %d \n", uint64(i), fact(uint64(i)))
	}
	end = time.Now()
	fmt.Printf("Calculation finished in %s \n", end.Sub(start)) //Calculation finished in 1ms

	fmt.Println("==================FACTORIAL MEMOIZATION==================")
	start = time.Now()
	var result uint64 = 0
	for i := uint64(0); i < LIM; i++ {
		result = FactorialMemoization(uint64(i))
		fmt.Printf("The factorial value for %d is %d\n", uint64(i), uint64(result))
	}

	end = time.Now()
	fmt.Printf("Calculation finished in %s\n", end.Sub(start)) // Calculation finished in 0ms

	fmt.Println("==================FACTORIAL FOR ==================")
	start = time.Now()

	fmt.Printf("The factorial 40 is %d\n", FactorialFor(40))

	end = time.Now()
	fmt.Printf("Calculation finished in %s\n", end.Sub(start)) // Calculation finished in 0ms

}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func FactorialFor(n uint64) (result uint64) {

	var j, x uint64

	if n == 0 {

		fmt.Println("Factorial closure for 0 is : 1")
		return 1
	}

	if n > 1 {

		x = 1
		for j = 1; j <= n; j++ {

			x = x * j

			fmt.Printf("Factorial closure for %d is : %d \n", j, x)
		}

		return x
	}

	fmt.Println("Factorial closure for 1 is : 1")
	return 1
}

func FactorialClosure() func(n uint64) uint64 {
	var a, b uint64 = 1, 1
	return func(n uint64) uint64 {
		if n > 1 {
			a, b = b, uint64(n)*uint64(b)
		} else {
			return 1
		}

		return b
	}
}

func FactorialMemoization(n uint64) (res uint64) {

	if facts[n] != 0 {
		res = facts[n]
		return res
	}

	if n > 0 {
		res = n * FactorialMemoization(n-1)
		facts[n] = res
		return res
	}

	return 1
}

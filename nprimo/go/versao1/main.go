package main

import "fmt"
import "time"

func primos(numero int) bool {
	for divisor := 2; divisor <= numero/2; divisor++ {
		if numero%divisor == 0 {
			return false
		}
	}
	return numero != 1
}
func main() {
	start := time.Now()
	for i := 1; i < 100000; i++ {
		if primos(i) {
			fmt.Sprintf("%d", i)
		}
	}
	end := time.Now()
	finish := end.Sub(start)
	fmt.Println("time:", finish)
}

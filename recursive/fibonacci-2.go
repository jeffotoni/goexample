package main

import "fmt"

func fibo(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		//fmt.Println(x)
	}
	return x
}

func main() {
	var input int
	fmt.Print("Enter the Number:")
	fmt.Scanf("%d", &input)
	fmt.Println("Result:", fibo(input))
}

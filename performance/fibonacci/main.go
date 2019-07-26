package main

import "fmt"
import "./pkg/fib"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(fib.Fib(i))
	}
}

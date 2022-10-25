package main

import "fmt"

func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	incrementA := wrapper()
	incrementB := wrapper()
	fmt.Println("A:", incrementA())
	fmt.Println("A:", incrementA())
	fmt.Println("B:", incrementB())
	fmt.Println("B:", incrementB())
	fmt.Println("B:", incrementB())
}

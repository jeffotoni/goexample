package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {

	first := 0
	second := 1
	var myFn = func() int {
		result := first
		temp := first
		first = second
		second = temp + second

		return result
	}
	return myFn
}

func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}

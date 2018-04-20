package main

import "C"

//export DoubleIt
func DoubleIt(x int) int {
	return x * 3
}

func main() {}

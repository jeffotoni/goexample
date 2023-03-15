package main

import "C"

//export Soma
func Soma(a, b int) int {
	return a + b
}

func main() {}

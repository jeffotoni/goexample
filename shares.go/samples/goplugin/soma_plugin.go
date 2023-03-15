package main

import "C"

//export Soma
func Soma(a, b int) int {
	return a + b
}

// go build -buildmode=plugin -o soma_plugin.so soma_plugin.go
func main() {}

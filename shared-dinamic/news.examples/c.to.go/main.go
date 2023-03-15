package main

// #include "soma.h"
import "C"
import "fmt"

func main() {
	x := 10
	y := 20
	result := C.add(C.int(x), C.int(y))
	fmt.Printf("Resultado: %d\n", result)
}

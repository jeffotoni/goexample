package main

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lsoma
// #include "soma.h"
import "C"
import "fmt"

// go build -o gsoma main.go
func main() {
	x := 15
	y := 20
	result := C.Soma(C.int(x), C.int(y))
	fmt.Printf("Resultado: %d\n", result)
}

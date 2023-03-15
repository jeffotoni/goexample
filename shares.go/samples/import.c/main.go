package main

/*
int soma(int a,int b) {
	return a+b;
}
*/
import "C"

import (
	"fmt"
)

func main() {
	r := C.soma(2, 10)
	fmt.Printf("resultado= %v\r\n", r)

}

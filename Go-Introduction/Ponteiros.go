package main

import "fmt"

func main() {
	//Ponteiros


	z := 10
	w := &z
	fmt.Println("O endereço do ponteiro é", &z)
	fmt.Println("O endereço do ponteiro é", w)
	fmt.Println("O endereço do ponteiro é", *w)
}

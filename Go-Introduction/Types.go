package main

import "fmt"

func main() {

	type cdss int
	var b cdss = 100

	fmt.Println(b)

	//conversao

	var x int
	x = int(b)

	fmt.Println(x)
}
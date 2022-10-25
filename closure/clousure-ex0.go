package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "Olha como ocorre as chamadas das funcoes."
		fmt.Println(y)
	}
	// fmt.Println(y) // outside scope of y
}

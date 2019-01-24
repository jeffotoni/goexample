// Go in action
// @jeffotoni
// 2019-01-16

package main

import "fmt"

func main() {

	// var a []string // wrong

	// An array of 10 integers
	var a1 [10]int
	a1[0] = 10
	fmt.Println(a1)

	// An array of 3 strings
	var a2 [3]string
	a2[0] = "Jeff"
	a2[1] = "Lambda"
	fmt.Println(a2)
	fmt.Println(a2[0], a2[1])

	// An array of 3 strings
	var a3 [5]string
	fmt.Println(a3)
}

// Go in action
// @jeffotoni
// 2019-01-16

package main

import "fmt"

func main() {

	// Letting Go compiler infer the length of the array
	a := [...]string{"C", "C++", "B", "Fortran", "Lisp", "Pascal", "Assembly"}
	fmt.Println(a)

	// Letting Go compiler infer the length of the array
	a2 := [...]string{}
	// a2[0] = "" // error invalid array index 0
	fmt.Println(a2)
	// a2 = append(a2, "Cloud") // error must be slice
	//
	a3 := [...]string{"@go_br", "@awsbrasil"}
	a3[0] = "Golang is life!"
	fmt.Println(a3)
}

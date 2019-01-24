// Go in action
// @jeffotoni
// 2019-01-16

package main

import "fmt"

func main() {

	// Declaring and initializing an array at the same time
	var a = [5]int{12, 24, 64, 55, 99}
	fmt.Println(a)

	// Declaring and initializing an array type string
	var a2 = [3]string{"lambda", "serverless", "Go"}
	fmt.Println(a2)

	// Short hand declaration
	a3 := [2]string{"2019", "Golang"}
	fmt.Println(a3)

	// Letting Go compiler infer the length of the array
	a4 := [...]int{34, 45, 57, 69, 611, 123, 174, 200, 223, 443, 445}
	fmt.Println(a4)

	// Letting Go compiler infer the length of the array
	a5 := [...]string{"C", "C++", "B", "Fortran", "Lisp", "Pascal", "Assembly"}
	fmt.Println(a5)
}

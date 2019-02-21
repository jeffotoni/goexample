// Go Api server
// @jeffotoni
// 2019-01-19

package main

import (
	"fmt"
)

// passed by reference
func setIntPtr(s **int, toadd *int) {
	*s = toadd
}

// passed by reference native
func add(p []int) {
	p = append(p, 56)
	fmt.Println(p)
}

func main() {

	///////////////////////////////////////////
	/// ex1
	var s *int
	fmt.Println(s)
	i := 3
	setIntPtr(&s, &i)
	fmt.Println(*s)
	fmt.Println("Hello, playground, ")

	///////////////////////////////////////////
	/// ex2
	person := make([]int, 1)
	fmt.Println(person)

	person[0] = 10
	fmt.Println(person)

	add(person)
	fmt.Println(person)
}

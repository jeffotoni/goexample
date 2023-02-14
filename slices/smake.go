// Go in action
// @jeffotoni
// 2019-01-16

package main

import "fmt"

func main() {
	a := make([]int, 4)
	a[0] = 12
	fmt.Println("a", a)

	b := make([]int, 0, 5)
	fmt.Println("b", b)

	c := b[:2]
	fmt.Println("c", c)
}

// Go in action
// @jeffotoni
// 2019-01-16

package main

import "fmt"

func main() {
	primes := [7]int{2, 3, 5, 7, 11, 13, 14}

	var p []int = primes[2:5]
	fmt.Println(p)
}

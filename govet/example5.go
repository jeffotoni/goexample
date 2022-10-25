/*
* Example make(chan)
*
* @package     main
* @author      @jeffotoni
* @size        19/02/2018
*
 */

package main

import "fmt"

func add(a int, b int) int {
	return a + b

	fmt.Println("unreachable")
	return 0
}

func div(a int, b int) int {
	if b == 0 {
		panic("division by 0")
	} else {
		return a / b
	}

	fmt.Println("unreachable")
	return 0
}

func fibonnaci(n int) int {
	switch n {
	case 0:
		return 1
	case 1:
		return 1
	default:
		return fibonnaci(n-1) + fibonnaci(n-2)
	}

	fmt.Println("unreachable")
	return 0
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(div(10, 2))
	fmt.Println(fibonnaci(5))
}

// wrong...
//...
package main

import (
	"fmt"
	"os"
	"strconv"
)

func Fibonacci(n int) int {

	if n == 0 {

		//fmt.Println(n)
		return 0

	} else if n == 1 {

		//fmt.Println(n)
		return 1

	} else {

		//fmt.Println(n)
		return (Fibonacci(n-1) + Fibonacci(n-2))
	}
}

func main() {

	if len(os.Args) == 2 {

		str := os.Args[1]
		n, _ := strconv.Atoi(str)

		fmt.Println("Numero de elementos:", n)

		fmt.Println(Fibonacci(n))

	} else {

		fmt.Println("Erro,o correto é: go run fibonacci 5")
	}

	//fmt.Println("vim-go")
}

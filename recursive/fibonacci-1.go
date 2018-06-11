package main

import (
	"fmt"
	"os"
	"strconv"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	tmp1, tmp2 := 0, 0
	return func() int {
		tmp1, tmp2 = tmp2, tmp1+tmp2
		if tmp2 == 0 {
			tmp1 = 1
		}
		return tmp2
	}
}

func main() {

	if len(os.Args) == 2 {

		str := os.Args[1]
		n, _ := strconv.Atoi(str)

		f := fibonacci()
		for i := 0; i < n; i++ {
			fmt.Println(f())

		}
	} else {

		fmt.Println("Erro, correct: go run fibonacci.go 10")
	}
}

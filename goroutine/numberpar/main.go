package main

import (
	"fmt"
	"runtime"
)

type T []int

func main() {
	runtime.GOMAXPROCS(2)
	c := make(chan int)
	i := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	go Par(i[0:10], c)
	go Par(i[10:],c )
	x, y :=  <-c, <-c 
	fmt.Println(x, y, " = ", x+y)
}

func Par(T []int, c chan int) {
	soma :=0
	for _, v := range T {
		if v%2 == 0 {
			soma += v
		}
	}
	c <- soma
}

package main

import (
	"fmt"
)

func goRoutineA(a chan<- int) {
	//val := <-a
	a <- 2
	fmt.Println("goRoutineA received the data")

}

func main() {

	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}

	ch := make(chan int)
	go goRoutineA(ch)

	x := <-ch
	fmt.Println("main goroutine..", x)
	//time.Sleep(time.Second * 2)
}

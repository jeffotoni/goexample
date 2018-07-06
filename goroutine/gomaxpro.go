package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	runtime.GOMAXPROCS(0)
	ch := make(chan int)
	n := 10000
	for i := 0; i < n; i++ {
		task(strconv.Itoa(i), ch, 100)
	}
	fmt.Printf("begin\n")
	for i := 0; i < n; i++ {
		<-ch
	}
}

func task(name string, ch chan int, max int) {
	go func() {
		i := 1
		for i <= max {
			time.Sleep(time.Second * 1)
			fmt.Printf("%s %d\n", name, i)
			//print(name + " " + strconv.Itoa(i) + "\n")
			i++
		}
		ch <- 1
	}()
}

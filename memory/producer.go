package main

import (
	"fmt"
	"time"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 900)
		chnl <- i
	}
	close(chnl)
}

func producer2(chnl chan int) {
	for i := 10; i < 20; i++ {
		time.Sleep(time.Second)
		chnl <- i
	}
	close(chnl)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	go producer2(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}

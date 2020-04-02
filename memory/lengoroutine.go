package main

import (
	"fmt"
	//"log"
	//"runtime"
	"time"
)

func main() {
	bs := make([]int, 1000000)
	doneCh := make(chan bool)
	stats()
	for i := 0; i < 1000000; i++ {
		go func(i int) {
			//fmt.Println("start:", i)
			<-time.After(1 * time.Second)
			//fmt.Println("finish:", i)
			bs = append(bs, i)
			fmt.Println(len(bs))
			//runtime.Ke																													epAlive(&bs)
			doneCh <- true
			return
		}(i)
	}
	//fmt.Println("done")

	<-doneCh
}

//go:noinline
func stats() {
	go func() {
		for {
			//print("NumGoroutine:")
			//log.Println(runtime.NumGoroutine())
			time.Sleep(1 * time.Second)
		}
	}()
}

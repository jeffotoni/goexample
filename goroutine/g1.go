package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	go say("world") // create a new goroutine
	say("hello")    // current goroutine
}

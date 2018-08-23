package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	type R struct {
		L sync.Mutex
	}

	var mutex = &sync.Mutex{}
	//var msg = make(chan string)
	var msgG bool

	// produz
	go func() {

		//msg <- "carrega string"
		mutex.Lock()
		msgG = true
		mutex.Unlock()

	}()

	// consome
	go func() {

		time.Sleep(time.Second * 2)
		mutex.Lock()
		x := msgG
		fmt.Println("Go1: ", x)
		mutex.Unlock()
		//x := <-msg
	}()

	// consome
	go func() {

		time.Sleep(time.Second * 1)
		//x := <-msg
		mutex.Lock()
		x := msgG
		fmt.Println("Go2: ", x)
		mutex.Unlock()
	}()

	//for {
	//	time.Sleep(time.Second * 5)
	//}

	time.Sleep(time.Second * 7)
	//fmt.Println(<-msg)
}

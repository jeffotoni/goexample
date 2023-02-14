package main

import (
	"fmt"
	"time"
)

func main() {

	max := 50
	running := make(chan bool, max)

	go func() {

		for range time.Tick(700 * time.Millisecond) {

			fmt.Printf("%d goroutines executando no momento\n", len(running))
		}
	}()

	go func() {
		for range time.Tick(200 * time.Millisecond) {

			fmt.Println("Goroutine disparada em 200 a 200 milli")
		}
	}()

	for i := 0; i < 10000; i++ {

		//time.Sleep(time.Millisecond * 10)
		running <- true
		fmt.Println("for::: running")

		//time.Sleep(time.Millisecond * 10)
		go func() {
			//fmt.Println("done2")
			defer func() {
				<-running
				//fmt.Println("done3")
				//time.Sleep(time.Second * 5)
			}()
		}()

		//time.Sleep(time.Millisecond * 50)
	}

	time.Sleep(time.Hour * 10)
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	_, done1 := doSomething1()
	//_, done2 := doSomething2()

	//do main thread

	println("wait for finish")
	<-done1
	//<-done2
	fmt.Print("clean up done, can exit safely")

}

func doSomething1() (error, chan bool) {
	//do something
	done := make(chan bool)
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTSTP)
	go func() {
		<-c
		//cleanup of something1
		done <- true
		fmt.Println("kill in your process!")
	}()
	return nil, done
}

func doSomething2() (error, chan bool) {
	//do something
	done := make(chan bool)
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		//cleanup of something2
		done <- true
	}()
	return nil, done
}

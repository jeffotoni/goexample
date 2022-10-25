package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	_, done1 := doSomething1()
	_, done2 := doSomething2()
	_, done3 := doSomething3()

	//do main thread

	println("wait for finish")
	<-done1
	<-done2
	<-done3

	fmt.Print("clean up done, can exit safely")

}

func doSomething3() (error, chan bool) {

	done := make(chan bool)
	c := make(chan os.Signal, 2)

	go func() {
		for {
			<-c
			<-time.After(time.Duration(5 * time.Second))
			println("for..")
			time.Sleep(time.Second * 3)
			done <- true
		}
	}()

	return nil, done
}

func doSomething1() (error, chan bool) {

	fmt.Println("dosomething1")
	//do something
	done := make(chan bool)
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		//cleanup of something1
		done <- true
	}()
	return nil, done
}

func doSomething2() (error, chan bool) {

	fmt.Println("dosomething2")
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

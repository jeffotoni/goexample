package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go ProcessSignal()

	done := make(chan bool, 1)

	fmt.Println("Open a new terminal, get the PID and issue kill -# PID command")

	// loop forever
	for {
		select {
		case <-done:
			break
		}

	}

}

func ProcessSignal() {

	sigch := make(chan os.Signal)

	// the case statement below will mute if not binded to signal.Notify
	// will purposely leave out SIGABRT(abort) as an example

	signal.Notify(sigch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGHUP, os.Interrupt)

	for {
		signalType := <-sigch
		fmt.Println("Received signal from channel : ", signalType)

		switch signalType {
		default:
			fmt.Printf("got signal = %v \n", signalType)
		case syscall.SIGHUP:
			fmt.Println("got Hangup/SIGHUP - portable number 1")
		case syscall.SIGINT:
			fmt.Println("got Terminal interrupt signal/SIGINT - portable number 2")
		case syscall.SIGQUIT:
			fmt.Println("got Terminal quit signal/SIGQUIT - portable number 3 - will core dump")
		case syscall.SIGABRT:
			fmt.Println("got Process abort signal/SIGABRT - portable number 6 - will core dump")
		case syscall.SIGKILL:
			fmt.Println("got Kill signal/SIGKILL - portable number 9")
		case syscall.SIGALRM:
			fmt.Println("got Alarm clock signal/SIGALRM - portable number 14")
		case syscall.SIGTERM:
			fmt.Println("got Termination signal/SIGTERM - portable number 15")
		case syscall.SIGUSR1:
			fmt.Println("got User-defined signal 1/SIGUSR1")
			//HINT : this is where you want to tell your program do something upon receiving a user-defined-signal

		case syscall.SIGUSR2:
			fmt.Println("got User-defined signal 2/SIGUSR2")

		}
	}
}

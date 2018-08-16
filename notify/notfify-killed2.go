package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("[SIGINT force quit]")
		os.Exit(0)
	}()
	err := Run()
	time.Sleep(time.Second * 20)
	fmt.Println(err)
}

func Run() error {
	cmd := exec.Command("ps")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

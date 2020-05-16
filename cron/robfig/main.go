package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("*/60 * * * * *", RunEverySecond)
	c.AddFunc("*/120 * * * * *", RunEveryTree)
	go c.Start()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig

}

func RunEverySecond() {
	fmt.Printf("%v\n", time.Now())
}

func RunEveryTree() {
	fmt.Println("tree:", time.Now())
}

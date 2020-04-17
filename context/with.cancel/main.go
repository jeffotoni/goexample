package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go f1(6, 1, cancel)
	go f1(2, 2, cancel)
	go f1(1, 3, cancel)

	select {
	case <-ctx.Done():
		fmt.Println("Gracefully exit")
		fmt.Println(ctx.Err())
		return
	case r := <-time.After(time.Duration(10) * time.Second):
		fmt.Println("f1():", r)
	}

	time.Sleep(time.Second * 10)
}

func f1(t, i int, cancel context.CancelFunc) {
	fmt.Println("processando goroutine:", i, " time: ", t)
	time.Sleep(time.Second * time.Duration(t))
	cancel()
	println("cancelando:", i)
}

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Body string
	sync.Mutex
}

func main() {

	var r Result
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go f1(6, 1, cancel, &r)
	go f1(2, 2, cancel, &r)
	go f1(1, 3, cancel, &r)

	select {
	case <-ctx.Done():
		fmt.Println("body:", r.Body)
		fmt.Println("Gracefully exit")
		fmt.Println(ctx.Err())
		return
	case r := <-time.After(time.Duration(10) * time.Second):
		fmt.Println("f1():", r)
	}

	time.Sleep(time.Second * 10)
}

func f1(t, i int, cancel context.CancelFunc, r *Result) {
	fmt.Println("processando goroutine:", i, " time: ", t)
	time.Sleep(time.Second * time.Duration(t))
	r.Lock()
	r.Body = `{"msg":"success"}`
	r.Unlock()
	cancel()
	println("cancelando:", i)
}

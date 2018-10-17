package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num goroutine 1:", runtime.NumGoroutine())

	for v := range gen(ctx) {
		fmt.Println(v)
		if v == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {

	dst := make(chan int)
	n := 1

	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			default:
				n++
				time.Sleep(time.Millisecond * 300)
				fmt.Println("working", n)
				//dst <- n
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num goroutine 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	// cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num goroutine 3:", runtime.NumGoroutine())

	return dst
}

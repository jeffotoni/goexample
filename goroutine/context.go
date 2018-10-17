package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.Background()

	fmt.Println("Context:\t", ctx)
	fmt.Println("Context:\t", ctx.Err())
	fmt.Printf("Context:\t%T\n", ctx)

	fmt.Println("-------------------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("Context:\t", ctx)
	fmt.Println("Context:\t", ctx.Err())
	fmt.Printf("Context:\t%T\n", ctx)
	fmt.Println("-------------------")

	cancel()

	fmt.Println("Context:\t", ctx)
	fmt.Println("Context:\t", ctx.Err())
	fmt.Println("Context:\t\t", cancel)
	fmt.Printf("Context:\t%T\n", cancel)
	fmt.Println("-------------------")

}

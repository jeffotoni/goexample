package main

import "fmt"

func main() {

	c := make(chan int, 2)
	d := make(chan int, 2)

	// basic
	go func() {

		d <- 10
		d <- 20
		c <- 30
		c <- <-d
	}()

	fmt.Println(<-d)
	//fmt.Println(<-d)
	fmt.Println(<-c)
	fmt.Println(<-c)

	close(d)
	close(c)

}

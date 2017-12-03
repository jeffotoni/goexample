/*
* Example parallelism
*
* @package     main
* @author      @jeffotoni
* @size        01/08/2017
*
 */

package main

import (
	"fmt"
	"time"
)

func EvenOne(even chan int) {
	i := 2
	for i < 9 {
		even <- i
		i = i + 2
	}
	close(even)
}

func OldOne(odd chan int) {
	i := 1
	for i < 9 {
		odd <- i
		i = i + 2

		time.Sleep(time.Second * 2)
	}
	close(odd)
}

func main() {

	even := make(chan int)
	odd := make(chan int)
	go EvenOne(even)
	go OldOne(odd)

	fmt.Println("for here:")

	for {

		even, ok1 := <-even
		odd, ok2 := <-odd

		if ok1 == false && ok2 == false {
			break
		}

		fmt.Println("Received even: ", even, ok1, " one: ", odd, ok2)
	}

	ch := make(chan int, 2)
	ch <- 5
	ch <- 6

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

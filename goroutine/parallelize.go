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
	"math/rand"
	"sync"
	"time"
)

// Parallelize parallelizes the function calls
func Paralle(functions ...func()) {

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(functions))

	defer waitGroup.Wait()

	for _, function := range functions {
		go func(copy func()) {
			defer waitGroup.Done()
			copy()
		}(function)
	}
}

func Dow(goroutine, loopindex int) {
	// simulate work
	time.Sleep(time.Second * time.Duration(5))
	fmt.Printf("gr[%d]: i=%d\n", goroutine, loopindex)
}

func f(n int) {
	for i := 0; i < 10; i++ {
		Dow(n, i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {

	func1 := func() {
		f(0)
	}

	func2 := func() {
		f(1)
	}

	func3 := func() {
		f(2)
	}

	Paralle(func1, func2, func3)
}

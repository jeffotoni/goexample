/*
* Example sync.Mutex
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
	//"strconv"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[int]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key int) {

	c.mux.Lock()

	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key] = key * 2

	//time.Sleep(time.Second * 1)
	//fmt.Println(c.v[key])

	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
//
//
func (c *SafeCounter) Value(key int) int {

	c.mux.Lock()

	// Lock so only one goroutine at a time can access the map c.v.
	//
	//
	defer c.mux.Unlock()

	return c.v[key]
}

func main() {

	c := SafeCounter{v: make(map[int]int)}

	for i := 0; i < 1000; i++ {

		go c.Inc(i)

	}

	time.Sleep(time.Millisecond * 300)

	for i := 0; i < 1000; i++ {

		time.Sleep(time.Millisecond * 300)
		ii := rand.Intn(999)

		Inc := c.Value(ii)

		if Inc == 0 {

			fmt.Println("Sleep 3 seconds")
			time.Sleep(time.Second * 1)
		}

		go fmt.Println("I:", ii, " value: ", Inc)
	}

	time.Sleep(time.Second * 1)

	// go func(key string) { fmt.Println(c.Value(key)) }("jeff")
	//fmt.Println(c.Value("jeff"))

}

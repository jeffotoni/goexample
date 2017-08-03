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
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()

	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
//
//
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()

	// Lock so only one goroutine at a time can access the map c.v.
	//
	//
	defer c.mux.Unlock()

	return c.v[key]
}

func main() {

	c := SafeCounter{v: make(map[string]int)}

	for i := 0; i < 1000; i++ {

		go c.Inc("jeff")

	}

	time.Sleep(time.Second * 1)

	// go func(key string) { fmt.Println(c.Value(key)) }("jeff")
	fmt.Println(c.Value("jeff"))

}

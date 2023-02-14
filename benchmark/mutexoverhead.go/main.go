package main

import (
	"fmt"
	"sync"
)

type unlocked struct {
	i int
}

type locked struct {
	i int
	l sync.Mutex
}

func (c *unlocked) add() {
	c.i++
}

func (c *locked) add() {
	c.l.Lock()
	defer c.l.Unlock()
	c.i++

	// c.l.Lock()
	// c.i++
	// c.l.Unlock()
}

func main() {
	fmt.Println("ok")
}

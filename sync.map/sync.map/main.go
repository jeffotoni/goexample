package main

import (
	"fmt"
	"sync"
)

type MyStruct struct {
	ID   int    `json:"id"`
	User string `json:"user"`
	Data string `json:"data"`
}

var sm sync.Map

func main() {
	var my MyStruct
	my.User = "jeffotoni"
	my.Data = "my data here... All Things"
	sm.Store(1, my)

	q, ok := sm.Load(1)
	fmt.Println(ok)
	fmt.Println(q)

	// cast
	mys := q.(MyStruct)
	fmt.Println(mys.User)
	fmt.Println(mys.Data)
}

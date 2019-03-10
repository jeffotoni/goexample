// Go in Action
// @jeffotoni
// 2019-03-10

package main

import (
	"container/list"
	"fmt"
)

func main() {

	// Create a new list and put some numbers in it.
	l := list.New()
	l.PushBack("jef1")
	l.PushBack("jef2")
	l.PushBack("jef3")
	l.PushBack("jef4")

	fmt.Println(l.Len())

	for l.Len() > 0 {
		e := l.Front() // First element
		fmt.Println(e.Value)
		l.Remove(e) // Dequeue
	}

	// not list .
	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("result1::", e.Value)
		//l.Remove(e)
	}
}

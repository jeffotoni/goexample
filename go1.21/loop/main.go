package main

import (
	"fmt"
	"sync"
)

func main() {
	myStuff := []string{"C", "C++", "Go"}
	wg := sync.WaitGroup{}
	for _, v := range myStuff {
		wg.Add(1)
		go func() {
			fmt.Println(v)
			wg.Done()
		}()
	}

	wg.Wait()

}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// var chars []string = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var chars []rune = []rune{'a'}

func main() {

	runtime.GOMAXPROCS(200)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Starting Go Routines")
	go func() {
		defer wg.Done()

		//i := 26

		for char := 'a'; char < 'a'+26; char++ {
			//for _, c := range chars {

			//char := int(c)
			fmt.Printf("%c", char)

		}
	}()

	go func() {
		defer wg.Done()

		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

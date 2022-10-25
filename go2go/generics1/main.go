package main

import (
	"fmt"
)

// The playground now supports parentheses or square brackets (only one at
// a time) for generic type and function declarations and instantiations.
// By default, parentheses are expected. To switch to square brackets,
// the first generic declaration in the source must use square brackets.

func Print[T any] (s []T){
	for _, v := range s {
		fmt.Print(v)

	}
	fmt.Println("")
}

func Reverse[T any](list []T) {
	i := 0
	j := len(list) - 1
	for i < j {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}
}

func main() {
	s := []string{"otoni", "jeff", "Go"}
	Reverse(s)
	Print(s)
	
	i := []int{2020, 2012, 2009}
	Reverse(i)
	Print(i)
}


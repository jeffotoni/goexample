package main

import (
	"fmt"
	"strconv"
)

type mapFunc func(int) int

type mapFunc2 func(int, string) string

type vet []int
type vet2 []string

func (s vet) Map1(mapper mapFunc) []int {

	var ret []int

	for _, v := range s {

		ret = append(ret, mapper(v))
	}

	return ret
}

func (s vet2) Map2(mapper mapFunc2) []string {

	var ret []string

	for i, v := range s {

		ret = append(ret, mapper(i, v))
	}

	return ret
}

func main() {

	fmt.Println("result: ", vet{1, 2, 3, 4}.Map1(func(x int) int { return x * 2 }))

	fmt.Println("result: ", vet2{"a", "b", "c"}.Map2(func(i int, x string) string { return strconv.Itoa(i) + " - " + x }))
}

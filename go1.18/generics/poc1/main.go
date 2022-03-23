package main

import "fmt"

type iMyinterface interface {
	int | int64 | int32
	~float64
	~string
	~[]byte
	MyMethod()
}

func NoGenericFuncInts(i []int) {
	for _, v := range i {
		fmt.Sprintf("%v", v)
	}
}

func NoGenericFuncStrs(s []string) {
	for _, v := range s {
		fmt.Sprintf("%v", v)
	}
}

func GenericsSlice[T any](s []T) {
	for _, v := range s {
		fmt.Sprintf("%v", v)
	}
}

func NoGenericInterface(i interface{}) {
	switch i.(type) {
	case []int:
		for _, v := range i.([]int) {
			fmt.Sprintf("%v", v)
		}
	case []string:
		for _, v := range i.([]string) {
			fmt.Sprintf("%v", v)
		}
	}
}

var any1 string
var T string

func main() {
	println("version...")
	GenericsSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	println("")
	GenericsSlice([]string{"a", "b", "c", "d"})
	println("")
	GenericsSlice[string]([]string{"a", "b", "c", "d", "j"})
	println("")
	GenericsSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	println("")
	NoGenericFuncInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	println("")
	NoGenericInterface([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	println("")
	NoGenericInterface([]string{"a", "b", "c", "d", "j"})

}

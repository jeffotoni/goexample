package main

import "fmt"

type S struct {
	x int
}

func F1(y int) *S {
	s := new(S)
	s.x = y
	s.x = s.x * 4
	return s
}

func F2(y int) *S {
	s := S{x: y}
	return &s
}

func F3() *int {
	y := 2
	res := y * 3
	return &res
}

func triInt(x int) int {
	x *= 3
	return x
}

func triArray(x [5]int) [5]int {
	for i := range x {
		x[i] *= 3
	}
	return x
}

func triSliceUpdate(x []int) []int {
	for i := range x {
		x[i] *= 3
	}
	return x
}

func triIntUpdate(x *int) int {
	*x *= 3
	return *x
}

func triArrayUpdate(x *[5]int) [5]int {
	for i := range x {
		x[i] *= 3
	}
	return *x
}

func triIntVariadic(x ...int) []int {
	for i := range x {
		x[i] *= 3
	}
	return x
}

func main() {
	x1 := 2
	fmt.Printf("triInt\t\t%v\t", x1)
	fmt.Printf("%v\t%v\n", triInt(x1), x1)
	x2 := [...]int{10, 20, 30, 40, 50}
	fmt.Printf("triArray\t%v\t", x2)
	fmt.Printf("%v\t%v\n", triArray(x2), x2)
	x3 := []int{10, 20, 30, 40, 50}
	fmt.Printf("triSliceUpdate\t%v\t", x3)
	fmt.Printf("%v\t%v\n", triSliceUpdate(x3), x3)
	x4 := 4
	x4ptr := &x4
	fmt.Printf("triIntUpdate\t%v\t", x4)
	fmt.Printf("%v\t%v\n", triIntUpdate(x4ptr), x4)

	x5 := [...]int{10, 20, 30, 40, 50}
	x5ptr := &x5
	fmt.Printf("triArrayUpdate\t%v\t", x5)
	fmt.Printf("%v\t%v\n", triArrayUpdate(x5ptr), x5)

	x6, x7, x8 := 100, 200, 300
	fmt.Printf("triIntVariadic\t%v, %v, %v\t", x6, x7, x8)
	fmt.Printf("%v\t%v, %v, %v\n", triIntVariadic(x6, x7, x8), x6, x7, x8)
	x9 := []int{10, 20, 30, 40, 50}
	fmt.Printf("triIntVariadic\t%v\t", x9)
	fmt.Printf("%v\t%v\n", triIntVariadic(x9...), x9)

	fmt.Printf("%v\t%v\n", F3(), F3())
}

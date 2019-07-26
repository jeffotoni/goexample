package main

import "fmt"

type User struct {
	Name string
	Uuid string
}

func main() {

	type slice struct {
		// pointer to underlying data in the slice.
		data uintptr
		// the number of elements in the slice.
		len int
		// the number of elements that the slice can
		// grow to before a new underlying array
		// is allocated.
		cap int
	}

	work1 := make([]*int, 1e5)

	work2 := make([]int, 15e8)

	vals := make([]string, 5, 5)

	fmt.Println(vals)

	vals2 := make([]string, 5)

	fmt.Println(vals2)

}

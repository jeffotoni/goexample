package main

import (
	"time"
)

type Item struct {
	Id   int
	time time.Time // aloca no heap
}
type Queue struct {
	B       [][]Item
	current time.Time // heap
	off     int
}

func main() {
	// ptr := new(int)
	// var x int = 42
	// var ptr *int = &x
	// ptr = &x
	// var ptr = &x
	// *ptr = 43
	// fmt.Println(*ptr)

	// x := [6]int{2, 3, 4, 5}
	// var s []int = x[1:4]
	// s[0] = 1

	//var a [2]string
	//a[0] = "jeff"
	// fmt.Println(a)
	// fmt.Sprintf("%v", a)

	//go func() {
	//	for {

	//	}
	//}()
	/*var c = make(chan bool)

	var a int
	var s []int // scape heap
	s = append(s, a)

	var t time.Time

	go func() {

		for {
			x := make([]int, 15e8)
			t = time.Now()
			x[0] = 1
		}
	}()

	<-c*/

	b := make([]int, 5, 5)
	b = append(b, 1)
	b = append(b, 2)
	b = append(b, 5)
	//fmt.Println(b, len(b), cap(b))

}

package main

func main() {
	u := make([]int, 8191) // Does not escape to heap
	_ = u

	v := make([]int, 8192) // Escapes to heap = 64kb
	_ = v

	var w [1024 * 1024 * 1.25]int
	_ = w

	var x [1024*1024*1.25 + 1]int // moved to heap > 10 mb
	_ = x
}

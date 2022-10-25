package main

type myNumber struct {
	n int
}

func plusOne(number *myNumber) {
	number.n++
}

func main() {
	var number *myNumber
	plusOne(number)
}

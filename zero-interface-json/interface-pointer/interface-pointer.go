package main

import "fmt"

type myNumber struct {
	n int
}

//func plusOne(number *myNumber) {
func plusOne(number *myNumber) {

	fmt.Println(number)

	if number != nil {
		number.n++
	}

	//number.n++

	fmt.Println("soma + 1: ", number)
}

func main() {

	//var number *myNumber
	var number = &myNumber{n: 10}
	plusOne(number)
}

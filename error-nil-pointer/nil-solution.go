package main

import "errors"

type myNumber struct {
	n int
}

/*
func TestNil() {
	var number *myNumber
	plusOne(number) // compile error: cannot use *myNumber as &myNumber, *myNumber can be nil
}

func TestPointer() {
	var number *myNumber = &myNumber{n: 5}
	plusOne(number) // compile error: cannot use *myNumber as &myNumber, *myNumber can be nil
}

func TestNonNilablePointer() {
	var number &myNumber = &myNumber{n: 5}
	plusOne(number)
	fmt.Println(number.n) // output: 6
}

func plusOnePointer(numberPointer *myNumber) error {
	if numberPointer == nil {
		return errors.New("number shouldn't be nil")
	}

	number := numberPointer.(*myNumber)
	plusOne(number)

}

func TestCastedPointer() {
	var number *myNumber = &myNumber{n: 5}

	plusOnePointer(number) // should handle error here
	fmt.Println(number.n)  // output: 6
}

*/

func plusOne(number *myNumber) error {

	if number == nil {
		return errors.New("number shouldn't be nil")
	}
	number.n++

	return nil
}

func main() {
	var number *myNumber
	plusOne(number)

	//number2 := number.(*myNumber)
	//plusOne(number2)
}

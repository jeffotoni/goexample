package main

import "fmt"

type myError struct{}

func main() {

	err1 := &myError{}
	err2 := fmt.Errorf("%w", err1)

	fmt.Println(err1)
	fmt.Println(err2)

	// errors.Unwrap(err2) == err1
	// errors.Is(err2, err1) == true

	// var me myError
	// errors.As(err2, &me) == true

	// fmt.Println(me)
}

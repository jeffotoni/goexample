package main

import "fmt"

func main() {

	mySlice := []bool{true, true, true}
	fmt.Printf("mySlice (before): %v \n", mySlice)
	clear(mySlice)
	fmt.Printf("mySlice (after): %v \n", mySlice)

	myMap := map[string]bool{"one": true, "two": true, "three": true}
	fmt.Printf("myMap (before): %v \n", myMap)
	clear(myMap)
	fmt.Printf("myMap (after): %v \n", myMap)

}

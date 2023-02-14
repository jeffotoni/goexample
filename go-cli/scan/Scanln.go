// Go in action
// @jeffotoni
// 2019-01-24

package main

import "fmt"

var name string
var age int

func main() {

	var input string

	fmt.Println("How many light years are we of the Sun?")

	fmt.Scanln(&input) //<--- here

	fmt.Println("Your answer : ", input)
	if input == "1,58" {
		fmt.Println("Congratulations is correct!")
	} else {
		fmt.Println("Correct answer is : 1,58 light years")
	}
}

// Go in action
// @jeffotoni
// 2019-01-24

package main

import (
	"fmt"
	"os"
)

var name string
var age int

func main() {

	fmt.Println("Your name:")
	if _, err := fmt.Scan(&name); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Your Age:")
	if _, err := fmt.Scan(&age); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Your name is: %s\n", name)
	fmt.Printf("Your age is: %d\n", age)
}

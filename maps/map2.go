// Go in action
// @jeffotoni
// 2019-01-24

package main

import "fmt"

func main() {

	// Required to initialize
	// the map with values
	var m = make(map[string]int)
	fmt.Println(m)
	if m == nil {
		fmt.Println("is nil")
	}
	m["population"] = 500000
	fmt.Println(m)
}

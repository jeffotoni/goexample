// Go in action
// @jeffotoni
// 2019-01-24

package main

import "fmt"

func main() {

	// Required to initialize
	// the map with values
	var m1 map[string]int
	var m2 = make(map[string]int)
	var m3 = map[string]int{"population": 500000}
	var m4 = map[string]int{"population": 500000}
	var m5 = m4
	var m6 map[string]string
	/* create a map*/
	m6 = make(map[string]string)
	fmt.Println(m1, m2, m3, m5, m6)
}

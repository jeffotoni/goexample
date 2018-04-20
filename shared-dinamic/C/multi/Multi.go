// @autor jeffotoni@gmail.com
package main

import "C"

//export Multi
func Multi(x int) int {
	return x * 2
}

func main() {}

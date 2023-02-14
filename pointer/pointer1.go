// Go in action
// @jeffotoni
// 2019-01-16

package main

func main() {

	var a int
	inc := &a
	*inc = 2
	*inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}

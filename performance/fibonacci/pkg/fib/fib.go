package fib

import "strings"

//import "strconv"

var x string
var b = strings.Builder{}

func Fib(n int) int {
	//x = x + "1"
	//x = string(n)
	b.Grow(50)
	b.WriteString(string(n))

	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

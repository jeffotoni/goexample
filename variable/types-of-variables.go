// Go in action
// @jeffotoni
// 2019-01-16

package main

import "fmt"

type T struct{ y int }

var x interface{} // x is nil and has static type interface{}
var v *T          // v has value nil, static type *T

type tv T

var (
	Float32    float32
	Float64    float64
	Boolean    bool
	Int        int
	String     = "@jeffotoni"
	Byte       = []byte("string here")
	Uint8      uint8
	Rune       rune
	Complex128 complex128
)

func main() {

	x = 42 // x has value 42 and dynamic type int
	fmt.Println(x)

	x = v // x has value (*T)(nil) and dynamic type *T
	fmt.Println(x)

	x = T{y: 2}
	fmt.Println(x)

	vx := tv{y: 10}
	fmt.Println(vx)
}

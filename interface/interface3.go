// Go in action
// @jeffotoni
// 2019-01-24

package main

import (
	"fmt"
)

type R struct {
	R string
}

type Iread interface {
	Read() string
}

func (r *R) Read() string {
	return fmt.Sprintf("Only: call Read")
}

func Call(ir Iread) string {
	return fmt.Sprintf("Read: %s", ir.Read())
}

func main() {
	var iread Iread
	r := R{"hello interface"}

	iread = &r
	fmt.Println(iread, r)
	fmt.Println(iread.Read())

	r2 := R{"hello interface call"}
	fmt.Println(Call(&r2))
}

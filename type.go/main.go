// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
)

type CName struct {
	Nome string
}

type HCName CName
type HI int
type HS string

func main() {
	var ii = HI(10)
	var ss = HS("otoni")
	var s = HCName{Nome: "Golang"}
	//s.Nome = "jefferson"
	fmt.Println("ok")
	Insert(CName(s))

	fmt.Println(reflect.TypeOf(CName(s)))
	fmt.Println(reflect.ValueOf(CName(s)).Kind())
	fmt.Println(reflect.ValueOf(s).Kind())
	fmt.Println(ii)
	fmt.Println(ss)
}

func Insert(c CName) {
	println("insert:", c.Nome)
}

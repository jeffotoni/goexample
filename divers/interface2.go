//
//
//
package main

import (
	"fmt"
)

type Inter1 interface {
	Metodh1() string
}

type Xmethod struct {
	Text string
}

func (x Xmethod) Metodh1() string {
	return x.Text
}

func TNext(s string) Xmethod {
	return Xmethod{s}
}

type YMethod struct {
	Xmethod
}

func TNewY() YMethod {

	return YMethod{TNext("...Hello World.. Folks...")}
}

func main() {

	var inter1 Inter1 = TNewY()
	fmt.Println(inter1.Metodh1())

	tnex := TNext("Hello brazil!!! ..Gophers!")
	fmt.Println(tnex.Metodh1())

}

/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size		   2017
 */

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

/////
///
///
type Cinterface interface {
	ShowC() *string
}

type Conf struct {
	Name string
	Cpf  string
}

func (cc *Conf) ShowC() *string {

	//fmt.Println("opa ShowC1()")
	cc.Name = "Jeff show c just only"
	return &cc.Name
}

func (cc *Conf) ShowC1() *string {

	//fmt.Println("opa ShowC1()")
	cc.Name = "Jeff show c"
	return &cc.Name
}

func (cc *Conf) ShowC2() string {

	//fmt.Println("opa ShowC()")
	cc.Name = "Jeff showc2"
	return cc.Name
}

func (cc *Conf) ShowC3() *Conf {

	cc.Name = "jef showc3 test!!"

	return cc
}

func ShowC4() *Conf {

	//return &Conf{s}
	return &Conf{

		Name: "jeff otoni show 4",
		Cpf:  "393.343.3343-48",
	}
}

func main() {

	var xc Conf

	fmt.Println(*xc.ShowC())
	fmt.Println(*xc.ShowC1())
	fmt.Println(xc.ShowC2())

	fmt.Println(xc.ShowC3().Name)

	fmt.Println("show4: ", ShowC4())
	fmt.Println("show4: ", ShowC4().Name)
	fmt.Println("show4: ", ShowC4().Cpf)

	var inter1 Inter1 = TNewY()
	fmt.Println(inter1.Metodh1())

	tnex := TNext("Hello brazil!!! ..Gophers!")
	fmt.Println(tnex.Metodh1())
}

package main

import (
	"fmt"
	"strconv"
)

type integer interface {
	Int() (int, error)
}

type mystr string

func (s mystr) Int() (int, error) {
	return strconv.Atoi(string(s))
}

func main() {
	var iface integer
	str := "42"

	iface = mystr(str)
	fmt.Println(iface.Int())
}


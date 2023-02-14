package main

import (
	"fmt"
)

func foo(s string, i int) (bool, error) {
	return false, nil
}

func bar(opts []string) (int, error) {
	return 10987, nil
}

var Foo func(string, int) (bool, error)
var Bar func([]string) (int, error)

func init() {
	Foo = foo
	Bar = bar
}

func main() {

	b, err := Foo("abc", 345)
	i, err := Bar([]string{"abc", "edf"})

	fmt.Println(b, err)
	fmt.Println(i, err)
}

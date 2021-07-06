package main

import "fmt"

type Options struct {
	Foo string
	Bar string
	Baz string
	Faz int
}

var gs int

func TakePointer(o *Options) []int {
	if o.Foo != "" {
		gs++
	}
	if o.Bar == "bar" {
		gs++
	}
	if o.Baz == "" {
		gs++
	}
	if o.Faz < 1_000 {
		gs++
	}

	var s []int
	s = append(s, gs)
	return s
}

func TakeCopy(o Options) {
	if o.Foo != "" {
		gs++
	}
	if o.Bar == "bar" {
		gs++
	}
	if o.Baz == "" {
		gs++
	}
	if o.Faz < 1_000 {
		gs++
	}
}

func DoublePassPointer(o *Options) {
	if o.Bar != "" {
		gs++
	}
	TakePointer(o)
}

func DoublePassCopy(o Options) {
	if o.Bar != "" {
		gs++
	}
	TakeCopy(o)
}

func makeOptionsPtr(foo, bar, baz string, faz int) *Options {
	return &Options{Foo: foo, Bar: bar, Baz: baz, Faz: faz}
}

func main() {
	fmt.Println("testando....")
}

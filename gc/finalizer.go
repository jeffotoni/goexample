package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

type Foo struct {
	a int
}

func main() {
	for i := 0; i < 1000; i++ {
		f := NewFoo(i)
		println(f.a)
		time.Sleep(time.Second)
	}

	runtime.GC()
}

//go:noinline
func NewFoo(i int) *Foo {
	f := &Foo{a: rand.Intn(1000)}
	runtime.SetFinalizer(f, func(f *Foo) {
		fmt.Println(`foo ` + strconv.Itoa(i) + ` has been garbage collected`)
	})

	return f
}

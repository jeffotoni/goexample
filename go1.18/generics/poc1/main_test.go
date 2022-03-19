package main

import (
	"fmt"
	"testing"
)

var tableInt = []struct {
	ints []int
}{
	{ints: []int{1, 2, 34, 4, 5, 6}},
}

var tableStr = []struct {
	strs []string
}{
	{strs: []string{"a", "b", "c", "d", "f", "g"}},
}

func BenchmarkGenericsA(b *testing.B) {
	for _, v := range tableInt {
		b.Run(fmt.Sprintf("%s", ""), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GenericsSlice(v.ints)
			}
		})
	}
}

func BenchmarkGenericsB(b *testing.B) {
	for _, v := range tableStr {
		b.Run(fmt.Sprintf("%s", ""), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GenericsSlice(v.strs)
			}
		})
	}
}

func BenchmarkGenericsF(b *testing.B) {
	for _, v := range tableInt {
		b.Run(fmt.Sprintf("%s", ""), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NoGenericFuncInts(v.ints)
			}
		})
	}
}

func BenchmarkGenericsS(b *testing.B) {
	for _, v := range tableStr {
		b.Run(fmt.Sprintf("%s", ""), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NoGenericFuncStrs(v.strs)
			}
		})
	}
}

func BenchmarkGenericsI(b *testing.B) {
	for _, v := range tableInt {
		b.Run(fmt.Sprintf("%s", ""), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NoGenericInterface(v.ints)
			}
		})
	}
}

func BenchmarkGenericsY(b *testing.B) {
	for _, v := range tableStr {
		b.Run(fmt.Sprintf("%s", ""), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NoGenericInterface(v.strs)
			}
		})
	}
}

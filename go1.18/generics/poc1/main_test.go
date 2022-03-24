package main

import (
	"fmt"
	"testing"
)

var tableInt = []struct {
	ints []int
}{
	{ints: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
}

var tableStr = []struct {
	strs []string
}{
	{strs: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "jeff", "otoni", "go", "C"}},
}

func BenchmarkNoGenericFuncInts___________(b *testing.B) {
	for _, v := range tableInt {
		b.Run(fmt.Sprintf("%s", ""), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NoGenericFuncInts(v.ints, v.ints)
			}
		})
	}
}

func BenchmarkNoGenericFuncStrs___________(b *testing.B) {
	for _, v := range tableStr {
		b.Run(fmt.Sprintf("%s", ""), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NoGenericFuncStrs(v.strs, v.strs)
			}
		})
	}
}

func BenchmarkGenericsComparableSliceInt(b *testing.B) {
	for _, v := range tableInt {
		b.Run(fmt.Sprintf("%s", ""), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GenericsComparableSlice(v.ints, v.ints)
			}
		})
	}
}

func BenchmarkGenericsComparableSliceStr(b *testing.B) {
	for _, v := range tableStr {
		b.Run(fmt.Sprintf("%s", ""), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GenericsComparableSlice(v.strs, v.strs)
			}
		})
	}
}

func BenchmarkNoGenericInterfaceInt(b *testing.B) {
	for _, v := range tableInt {
		b.Run(fmt.Sprintf("%s", ""), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NoGenericInterface(v.ints, v.ints)
			}
		})
	}
}

func BenchmarkNoGenericInterfaceStr(b *testing.B) {
	for _, v := range tableStr {
		b.Run(fmt.Sprintf("%s", ""), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NoGenericInterface(v.strs, v.strs)
			}
		})
	}
}

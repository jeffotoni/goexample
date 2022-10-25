package main

import (
	"fmt"
	"strconv"
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

var m = make(map[string]int)
var m1 = make(map[string]int)
var m2 = make(map[int]int)

func BenchmarkMapNoGenericDelete____________(b *testing.B) {
	for ii := 0; ii < 10000; ii++ {
		pos := strconv.Itoa(ii)
		m[pos] = ii * 2
	}
	for ii1 := 0; ii1 < 10000; ii1++ {
		pos := strconv.Itoa(ii1)
		m1[pos] = ii1 * 2
	}
	for i := 0; i < b.N; i++ {
		deleteMap(m)
		deleteMap(m1)
	}
}
func BenchmarkMapGenericDelete______________(b *testing.B) {
	for ii := 0; ii < 10000; ii++ {
		pos := strconv.Itoa(ii)
		m[pos] = ii * 2
	}
	for ii2 := 0; ii2 < 10000; ii2++ {
		m2[ii2] = ii2 * 3
	}
	for i := 0; i < b.N; i++ {
		gdeleteMap(m)
		gdeleteMap(m2)
	}
}

func BenchmarkNoGenericFuncInts_________(b *testing.B) {
	for _, v := range tableInt {
		b.Run(fmt.Sprintf("%s", ""), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NoGenericFuncInts(v.ints, v.ints)
			}
		})
	}
}

func BenchmarkNoGenericFuncStrs_________(b *testing.B) {
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

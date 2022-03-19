package main

import (
	"fmt"
	"testing"
)

var table = []struct {
	ints   int64
	floats float64
}{
	{ints: 100, floats: 44.55},
	{ints: 1000, floats: 55.66},
	{ints: 74382, floats: 67.77},
	{ints: 382399, floats: 88.88},
}

func BenchmarkGenericsOrIntFloat(b *testing.B) {
	var mi = make(map[string]int64)
	var mf = make(map[string]float64)

	for _, v := range table {
		b.Run(fmt.Sprintf("ints_size_%v floats_size_%v", v.ints, v.floats), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mi["first"] = v.ints
				mf["second"] = v.floats
				GenericsOrIntFloat(mi, mf)
			}
		})
	}
}

func BenchmarkGenericsComparable(b *testing.B) {
	var mi = make(map[string]int64)
	var mf = make(map[string]float64)

	for _, v := range table {
		b.Run(fmt.Sprintf("ints_size_%v floats_size_%v", v.ints, v.floats), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mi["first"] = v.ints
				mf["second"] = v.floats
				GenericsComparable(mi, mf)
			}
		})
	}
}

func BenchmarkGenericsInferred(b *testing.B) {
	var mi = make(map[string]int64)
	var mf = make(map[string]float64)

	for _, v := range table {
		b.Run(fmt.Sprintf("ints_size_%v floats_size_%v", v.ints, v.floats), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mi["first"] = v.ints
				mf["second"] = v.floats
				GenericsInferred(mi, mf)
			}
		})
	}
}

func BenchmarkNoGenericsFuncs(b *testing.B) {
	var mi = make(map[string]int64)
	var mf = make(map[string]float64)

	for _, v := range table {
		b.Run(fmt.Sprintf("ints_size_%v floats_size_%v", v.ints, v.floats), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mi["first"] = v.ints
				mf["second"] = v.floats
				noGenericsFuncs(mi, mf)
			}
		})
	}
}

func BenchmarkNoGenericsInterface(b *testing.B) {
	var mi = make(map[string]int64)
	var mf = make(map[string]float64)

	for _, v := range table {
		b.Run(fmt.Sprintf("ints_size_%v floats_size_%v", v.ints, v.floats), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mi["first"] = v.ints
				mf["second"] = v.floats
				noGenericsInterface(mi, mf)
			}
		})
	}
}

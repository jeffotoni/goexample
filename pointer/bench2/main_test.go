package main

import (
	"testing"
)

func BenchmarkF1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var s int
		s++
		F1(s)
	}
}

func BenchmarkF2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var s int
		s++
		F2(s)
	}
}

func BenchmarkF3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var s int
		s++
		F3(s)
	}
}

func BenchmarkTriInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		x := 1
		triInt(x)
	}
}

func BenchmarkTriArray(b *testing.B) {
	for n := 0; n < b.N; n++ {
		x2 := [...]int{10, 20, 30, 40, 50}
		triArray(x2)
	}
}

func BenchmarkTriSliceUpdate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		x3 := []int{10, 20, 30, 40, 50}
		triSliceUpdate(x3)
	}
}

func BenchmarkTriIntUpdate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		x4 := 4
		x4ptr := &x4
		triIntUpdate(x4ptr)
	}
}

func BenchmarkTriArrayUpdate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		x5 := [...]int{10, 20, 30, 40, 50}
		x5ptr := &x5
		triArrayUpdate(x5ptr)
	}
}

func BenchmarkTriIntVariadic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		x6, x7, x8 := 100, 200, 300
		triIntVariadic(x6, x7, x8)
	}
}

func BenchmarkTriIntVariadic2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		x9 := []int{10, 20, 30, 40, 50}
		triIntVariadic(x9...)
	}
}

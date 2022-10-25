package main

import "testing"

var kb8 [1024]int
var kb8Pt *[1024]int

var b80 [10]int
var b80Pt *[10]int

var mb8 [1024 * 1024]int
var mb8Pt *[1024 * 1024]int

func BenchmarkCopying80b(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b80 = return80b()
	}
}

func BenchmarkCopying80bPointer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b80Pt = return80bPointer()
	}
}

func BenchmarkCopying8kb(b *testing.B) {
	for n := 0; n < b.N; n++ {
		kb8 = return8kb()
	}
}

func BenchmarkCopying8kbPointer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		kb8Pt = return8kbPointer()
	}
}

func BenchmarkCopying8mb(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mb8 = return8mb()
	}
}

func BenchmarkCopying8mbPointer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mb8Pt = return8mbPointer()
	}
}

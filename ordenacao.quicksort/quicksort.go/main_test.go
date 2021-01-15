package main

import (
	"testing"
)

func BenchmarkMakeGenerate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		generateSlice(i)
	}
}

func BenchmarkMakeGenerate2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		generateSlice2(i)
	}
}

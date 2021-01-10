package main

import (
	"testing"
)

func BenchmarkMakeGenerate(b *testing.B) {
	//big := NewBig()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//rand.Int()
		//big.Len()
		generateSlice(i)
	}
}

func BenchmarkMakeGenerate2(b *testing.B) {
	//big := NewBig()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//rand.Int()
		//big.Len()
		generateSlice2(i)
	}
}

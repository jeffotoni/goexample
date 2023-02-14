// main_test.go
package main

import (
	"testing"
)

var str string

func BenchmarkStringBuildNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str = Naive()
	}
}
func BenchmarkStringBuildBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str = Builder()
	}
}

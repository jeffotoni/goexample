package main

import "testing"

func BenchmarkMake(b *testing.B) {
  for n := 0; n < b.N; n++ {
    data := make([]int, 0)
    for k := 0; k < size; k++ {
      data = append(data, k)
    }
  }
}

const size = 1000000

func BenchmarkMakeSize(b *testing.B) {
  for n := 0; n < b.N; n++ {
    data := make([]int, 0, size)
    for k := 0; k < size; k++ {
      data = append(data, k)
    }
  }
}

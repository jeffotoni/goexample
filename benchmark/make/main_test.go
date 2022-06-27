package main

import "testing"

const size = 1000000

func BenchmarkMake(b *testing.B) {
  for n := 0; n < b.N; n++ {
    data := make([]int, 0)
    for i := 0; i < size; i++ {
      data = append(data, i)
    }
  }
}

func BenchmarkMakeNil(b *testing.B) {
  for n := 0; n < b.N; n++ {
    data := []int{}
    for i := 0; i < size; i++ {
      data = append(data, i)
    }
  }
}

func BenchmarkMakeSize(b *testing.B) {
  for n := 0; n < b.N; n++ {
    data := make([]int, 0, size)
    for i := 0; i < size; i++ {
      data = append(data, i)
    }
  }
}

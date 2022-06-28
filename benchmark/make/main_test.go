package main

import "testing"

const size = 1000000

func BenchmarkMakeAppend(b *testing.B) {
  for n := 0; n < b.N; n++ {
    data := make([]int, 0)
    for i := 0; i < size; i++ {
      data = append(data, i)
    }
  }
}

func BenchmarkWithoutMakeAppend(b *testing.B) {
  for n := 0; n < b.N; n++ {
    data := []int{}
    for i := 0; i < size; i++ {
      data = append(data, i)
    }
  }
}

func BenchmarkMakeNilWithoutAppend(b *testing.B) {
  for n := 0; n < b.N; n++ {
    data := [size]int{}
    for i := 0; i < size; i++ {
      data[i] = i * 3
    }
  }
}

func BenchmarkMakeSizeAppend(b *testing.B) {
  for n := 0; n < b.N; n++ {
    data := make([]int, 0, size)
    for i := 0; i < size; i++ {
      data = append(data, i)
    }
  }
}

func BenchmarkMakeMap(b *testing.B) {
  for n := 0; n < b.N; n++ {
    data := make(map[int]int, 0)
    for i := 0; i < size; i++ {
      data[i] = i * 2
    }
  }
}

func BenchmarkWithoutMakeMapNil(b *testing.B) {
  for n := 0; n < b.N; n++ {
    data := map[int]int{}
    for i := 0; i < size; i++ {
      data[i] = i * 2
    }
  }
}

func BenchmarkMakeMapInterface(b *testing.B) {
  for n := 0; n < b.N; n++ {
    data := make(map[int]interface{}, 0)
    for i := 0; i < size; i++ {
      data[i] = i * 2
    }
  }
}

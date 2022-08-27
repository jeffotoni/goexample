package main

import "testing"

func BenchmarkFixBirdCountLog(b *testing.B) {
  
  for n := 0; n < b.N; n++ {
    _ = FixBirdCountLog([]int{n+0, n+1, n+2, n+3, n+4})
  }
}

func BenchmarkFixBirdCountLog2(b *testing.B) {
  for n := 0; n < b.N; n++ {
  _ = FixBirdCountLog([]int{n+0, n+1, n+2, n+3, n+4})
  }
}
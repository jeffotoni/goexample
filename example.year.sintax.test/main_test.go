package main

import (
  "math/rand"
  "testing"
  "time"
)

func BenchmarkYearOne(b *testing.B) {
  for n := 0; n < b.N; n++ {
    year := randInt(1000, 5050)
    IsLeapYearOne(year)
  }
}

func IsLeapYearOne(year int) bool {
  return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func BenchmarkYearTwo(b *testing.B) {
  for n := 0; n < b.N; n++ {
    year := randInt(1000, 5050)
    IsLeapYearTwo(year)
  }
}

func IsLeapYearTwo(year int) bool {
  if year%4 == 0 {
    if year%100 == 0 {
      if year%400 == 0 {
        return true
      } else {
        return false
      }
    } else {
      return true
    }
  }
  return false
}

func BenchmarkYearThree(b *testing.B) {
  for n := 0; n < b.N; n++ {
    year := randInt(1000, 5050)
    IsLeapYearThree(year)
  }
}

func IsLeapYearThree(y int) bool {

  divisibleBy4 := y%4 == 0
  divisibleBy100 := y%100 == 0
  divisibleBy400 := y%400 == 0

  if divisibleBy4 {
    if divisibleBy100 {
      if divisibleBy400 {
        return true
      }
      return false
    }
    return true
  }

  return false
}

func randInt(min int, max int) int {
  rand.Seed(time.Now().UTC().UnixNano())
  return min + rand.Intn(max-min)
}

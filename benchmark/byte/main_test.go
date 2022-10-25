package main

import (
  "net/http/httptest"
  "testing"
)

func BenchmarkByte(b *testing.B) {
  var w = httptest.NewRecorder()
  for n := 0; n < b.N; n++ {
    w.Write([]byte("Hello Jeffotoni!"))
  }
}

func BenchmarkByteData(b *testing.B) {
  var w = httptest.NewRecorder()
  data := []byte("Hello Jeffotoni!")
  for n := 0; n < b.N; n++ {
    w.Write(data)
  }
}

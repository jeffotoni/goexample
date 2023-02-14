package main

import (
	"bytes"
	"fmt"
	"testing"
)

func BenchmarkString(b *testing.B) {
	bs := []byte{104, 101, 108, 108, 111}
	for i := 0; i < b.N; i++ {
		_ = string(bs)
	}
}

func BenchmarkSprintf(b *testing.B) {
	bs := []byte{104, 101, 108, 108, 111}
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s", bs)
	}
}

func BenchmarkBuffer(b *testing.B) {
	bs := []byte{104, 101, 108, 108, 111}
	for i := 0; i < b.N; i++ {
		_ = bytes.NewBuffer(bs).String()
	}
}

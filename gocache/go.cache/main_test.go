package main

import (
	"testing"
	"time"
)

func BenchmarkGoCacheSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tt, _ := time.ParseDuration("1800000ms")
		Set("key", n, tt)
	}
}

func BenchmarkGoCacheGet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Get("key")
	}
}

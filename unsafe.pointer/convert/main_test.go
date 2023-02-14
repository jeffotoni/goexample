package main

import "testing"

func BenchmarkUnsafeBytesToString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bs := []byte("Test")
		UnsafeBytesToString(bs)
	}
}

func BenchmarkSafeBytesToString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bs := []byte("Test")
		SafeBytesToString(bs)
	}
}

func BenchmarkSafeStringToBytes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bs := string("Test")
		SafeStringToBytes(bs)
	}
}

func BenchmarkUnsafeStringToBytes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bs := string("Test")
		UnsafeStringToBytes(bs)
	}
}

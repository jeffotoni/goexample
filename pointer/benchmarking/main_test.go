package main

import (
	"testing"
)

func BenchmarkPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TakePointer(&Options{Foo: "foo", Bar: "bar", Baz: "baz", Faz: 10})
	}
}

func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TakeCopy(Options{Foo: "foo", Bar: "bar", Baz: "baz", Faz: 10})
	}
}

func BenchmarkDoublePassPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoublePassPointer(&Options{Foo: "foo", Bar: "bar", Baz: "baz", Faz: 10})
	}
}

func BenchmarkDoublePassCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoublePassCopy(Options{Foo: "foo", Bar: "bar", Baz: "baz", Faz: 10})
	}
}

func BenchmarkReusePointer(b *testing.B) {
	o := &Options{Foo: "foo", Bar: "bar", Baz: "baz", Faz: 10}
	for i := 0; i < b.N; i++ {
		DoublePassPointer(o)
	}
}

func BenchmarkReuseCopy(b *testing.B) {
	o := Options{Foo: "foo", Bar: "bar", Baz: "baz", Faz: 10}
	for i := 0; i < b.N; i++ {
		DoublePassCopy(o)
	}
}

func BenchmarkReusePointerMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoublePassPointer(makeOptionsPtr("foo", "bar", "baz", 10))
	}
}

package main

import (
	"strings"
	"testing"
)

func CountString(s string, b byte) (c int) {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			c++
		}
	}
	return
}

func CountBytes(s []byte, b byte) (c int) {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			c++
		}
	}
	return
}

func CountGeneric[S string | []byte](s S, b byte) (c int) {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			c++
		}
	}
	return
}

func TestCount(t *testing.T) {
	for _, td := range []struct {
		input  string
		char   byte
		expect int
	}{
		{"", 'x', 0},
		{"xxxx", 'x', 4},
		{".....", 'x', 0},
		{".....x", 'x', 1},
	} {
		t.Run("", func(t *testing.T) {
			t.Run("bytes", func(t *testing.T) {
				if a := CountBytes([]byte(td.input), td.char); a != td.expect {
					t.Fatalf("expected: %d; actual: %d", td.expect, a)
				}
			})
			t.Run("str", func(t *testing.T) {
				if a := CountString(td.input, td.char); a != td.expect {
					t.Fatalf("expected: %d; actual: %d", td.expect, a)
				}
			})
			t.Run("generic_bytes", func(t *testing.T) {
				if a := CountGeneric([]byte(td.input), td.char); a != td.expect {
					t.Fatalf("expected: %d; actual: %d", td.expect, a)
				}
			})
			t.Run("generic_str", func(t *testing.T) {
				if a := CountGeneric(td.input, td.char); a != td.expect {
					t.Fatalf("expected: %d; actual: %d", td.expect, a)
				}
			})
		})
	}
}

var Gi int

func BenchmarkCount(b *testing.B) {
	for _, td := range []struct {
		name  string
		input string
		char  byte
	}{
		{"empty", "", 'x'},
		{"1kib_all", MakeStr("x", 1024), 'x'},
		{"1mib_all", MakeStr("x", 1024*1024), 'x'},
		{"1mib_nomatch", MakeStr("x", 1024*1024), 'y'},
	} {
		b.Run(td.name, func(b *testing.B) {
			bytes := []byte(td.input)
			b.ResetTimer()

			b.Run("bytes", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					Gi = CountBytes(bytes, td.char)
				}
			})
			b.Run("str", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					Gi = CountString(td.input, td.char)
				}
			})
			b.Run("generic_bytes", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					Gi = CountGeneric(bytes, td.char)
				}
			})
			b.Run("generic_str", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					Gi = CountGeneric(td.input, td.char)
				}
			})
		})
	}
}

func MakeStr(s string, n int) string {
	var b strings.Builder
	b.Grow(len(s) * n)
	for i := 0; i < n; i++ {
		b.WriteString(s)
	}
	return b.String()
}


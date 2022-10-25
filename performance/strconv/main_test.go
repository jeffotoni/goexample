// main_test.go
package main

import (
    "testing"
)

var (
    a    = "boo"
    blah = 42
    box  = ""
)

func BenchmarkStrconv(b *testing.B) {
    for i := 0; i < b.N; i++ {
        box = strconvFmt(a, blah)
    }
    a = box
}

func BenchmarkFmt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        box = fmtFmt(a, blah)
    }
    a = box
}

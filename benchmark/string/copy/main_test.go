package main

import (
    "bytes"

    "fmt"

    "strings"
    "testing"
)

func BenchmarkConcat(b *testing.B) {
    var str string
    for n := 0; n < b.N; n++ {
        str += "jeffotoni" + " Go is life!"
    }
    b.StopTimer()

    // if s := strings.Repeat("x", b.N); str != s {
    //     b.Errorf("unexpected result; got=%s, want=%s", str, s)
    // }
}

func BenchmarkBuffer(b *testing.B) {
    var buffer bytes.Buffer
    for n := 0; n < b.N; n++ {
        buffer.WriteString("jeffotoni")
        buffer.WriteString(" Go is life!")
    }
    b.StopTimer()
}

func BenchmarkCopy(b *testing.B) {
    bs := make([]byte, 0, b.N)
    bl := 0

    b.ResetTimer()
    for n := 0; n < b.N; n++ {
        bl += copy(bs[bl:], "jeffotoni")
        bl += copy(bs[bl:], " Go is life!")
    }
    b.StopTimer()
}

func BenchmarkStringBuilder(b *testing.B) {
    var strBuilder strings.Builder
    b.ResetTimer()
    for n := 0; n < b.N; n++ {
        strBuilder.WriteString("jeffotoni")
        strBuilder.WriteString(" Go is life!")
    }
    b.StopTimer()

    // if s := strings.Repeat("x", b.N); strBuilder.String() != s {
    //     b.Errorf("unexpected; got=%s, want=%s", strBuilder.String(), s)
    // }
}

func BenchmarkSprintf(b *testing.B) {
    var str1 = "jeffotoni"
    var str2 = " Go is life!"
    var out string

    b.ResetTimer()
    for n := 0; n < b.N; n++ {
        out = fmt.Sprintf("%s %s ", str1, str2)
    }
    b.StopTimer()

    if s := strings.Repeat("x", b.N); fmt.Sprintf("%s", s) != s {
        b.Errorf("unexpected result; got=%s, want=%s", fmt.Sprintf("%s %s", s, out), s)
    }
}

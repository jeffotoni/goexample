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
        str += "x"
    }
    b.StopTimer()

    if s := strings.Repeat("x", b.N); str != s {
        b.Errorf("unexpected result; got=%s, want=%s", str, s)
    }
}

func BenchmarkBuffer(b *testing.B) {
    var buffer bytes.Buffer
    for n := 0; n < b.N; n++ {
        buffer.WriteString("x")
    }
    b.StopTimer()

    if s := strings.Repeat("x", b.N); buffer.String() != s {
        b.Errorf("unexpected result; got=%s, want=%s", buffer.String(), s)
    }
}

func BenchmarkCopy(b *testing.B) {
    bs := make([]byte, b.N)
    bl := 0

    b.ResetTimer()
    for n := 0; n < b.N; n++ {
        bl += copy(bs[bl:], "x")
    }
    b.StopTimer()

    if s := strings.Repeat("x", b.N); string(bs) != s {
        b.Errorf("unexpected result; got=%s, want=%s", string(bs), s)
    }
}

func BenchmarkStringBuilder(b *testing.B) {
    var strBuilder strings.Builder
    b.ResetTimer()
    for n := 0; n < b.N; n++ {
        strBuilder.WriteString("x")
    }
    b.StopTimer()

    if s := strings.Repeat("x", b.N); strBuilder.String() != s {
        b.Errorf("unexpected; got=%s, want=%s", strBuilder.String(), s)
    }
}

func BenchmarkSprintf(b *testing.B) {
    var str1 = "string1"
    var str2 = "string2"
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

package main_test

import (
    "fmt"
    "strings"
    "testing"
)

var str, longStr string = "my_string", `qwertyuiopqwertyuiopqwertyuio
qwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiop`

const cStr = "my_string"

func BenchmarkPlus(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = "my_string" + str
    }
}

func BenchmarkLongPlus(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = "my_string" + longStr
    }
}

func BenchmarkConstPlus(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = "my_string" + cStr
    }
}

func BenchmarkJoin(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = strings.Join([]string{"my_string%s", str}, "")
    }
}

func BenchmarkLongJoin(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = strings.Join([]string{"my_string%s", longStr}, "")
    }
}

func BenchmarkConstJoin(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = strings.Join([]string{"my_string%s", cStr}, "")
    }
}
func BenchmarkSprintf(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = fmt.Sprintf("my_string%s", str)
    }
}

func BenchmarkLongSprintf(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = fmt.Sprintf("my_string%s", longStr)
    }
}

func BenchmarkConstSprintf(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _ = fmt.Sprintf("my_string%s", cStr)
    }
}

func BenchmarkBuilder(b *testing.B) {
    for n := 0; n < b.N; n++ {
        var b strings.Builder
        b.WriteString("my_string")
        b.WriteString(longStr)
    }
}

func BenchmarkLongBuilder(b *testing.B) {
    for n := 0; n < b.N; n++ {
        var b strings.Builder
        b.WriteString("my_string")
        b.WriteString(longStr)
    }
}

func BenchmarkConstBuilder(b *testing.B) {
    for n := 0; n < b.N; n++ {
        var b strings.Builder
        b.WriteString("my_string")
        b.WriteString(longStr)
    }
}

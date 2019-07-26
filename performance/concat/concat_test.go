package concat_test

import (
    "bytes"
    "testing"
)

const (
    concatSteps = 100
)

func BenchmarkConcat(b *testing.B) {
    for n := 0; n < b.N; n++ {
        var str string
        for i := 0; i < concatSteps; i++ {
            str += "x"
        }
    }
}

func BenchmarkBuffer(b *testing.B) {
    for n := 0; n < b.N; n++ {
        var buffer bytes.Buffer
        for i := 0; i < concatSteps; i++ {
            buffer.WriteString("x")
        }
    }
}

func BenchmarkAppend(b *testing.B) {
    //var xstr string
    var strs []string
    for n := 0; n < b.N; n++ {
        for i := 0; i < concatSteps; i++ {
            strs = append(strs, "x")
        }
    }
}

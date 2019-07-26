// main.go
package main

import "strings"

var strs = []string{
    "Jeffer",
    "Otoni",
    "@jeffotoni",
    "golang",
    "go",
    "c",
    "c++",
    "assembly",
    "Occam",
    "Oberon",
    "Limbo",
    "Alef",
    "lambda",
}

func Naive() string {
    var s string

    for _, v := range strs {
        s += v
    }

    return s
}

func Builder() string {
    b := strings.Builder{}

    // Grow the buffer to a decent length, so we don't have to continually
    // re-allocate.
    b.Grow(60)

    for _, v := range strs {
        b.WriteString(v)
    }

    return b.String()
}

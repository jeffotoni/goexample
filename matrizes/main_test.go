package main

import (
    "os"
    "runtime/trace"
    "testing"
)

type Dot func(B, C *Matrix) error

var (
    A = &Matrix{
        N: 8,
        data: [][]float64{
            {1, 2, 3, 4, 5, 6, 7, 8},
            {9, 1, 2, 3, 4, 5, 6, 7},
            {8, 9, 1, 2, 3, 4, 5, 6},
            {7, 8, 9, 1, 2, 3, 4, 5},
            {6, 7, 8, 9, 1, 2, 3, 4},
            {5, 6, 7, 8, 9, 1, 2, 3},
            {4, 5, 6, 7, 8, 9, 1, 2},
            {3, 4, 5, 6, 7, 8, 9, 0},
        },
    }
    B = &Matrix{
        N: 8,
        data: [][]float64{
            {9, 8, 7, 6, 5, 4, 3, 2},
            {1, 9, 8, 7, 6, 5, 4, 3},
            {2, 1, 9, 8, 7, 6, 5, 4},
            {3, 2, 1, 9, 8, 7, 6, 5},
            {4, 3, 2, 1, 9, 8, 7, 6},
            {5, 4, 3, 2, 1, 9, 8, 7},
            {6, 5, 4, 3, 2, 1, 9, 8},
            {7, 6, 5, 4, 3, 2, 1, 0},
        },
    }
    C = &Matrix{
        N: 8,
        data: [][]float64{
            {0, 0, 0, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0, 0, 0, 0},
            {0, 0, 0, 0, 0, 0, 0, 0},
        },
    }
)

func BenchmarkMatrixDotNaive(b *testing.B) {
    f, _ := os.Create("bench.trace")
    defer f.Close()
    trace.Start(f)
    defer trace.Stop()

    tests := []struct {
        name string
        f    Dot
    }{
        {
            name: "A.MultNaive",
            f:    A.MultNaive,
        },
        {
            name: "A.ParalMultNaivePerRow",
            f:    A.ParalMultNaivePerRow,
        },
        {
            name: "A.ParalMultNaivePerElem",
            f:    A.ParalMultNaivePerElem,
        },
    }
    for _, tt := range tests {
        b.Run(tt.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                tt.f(B, C)
            }
        })
    }
}

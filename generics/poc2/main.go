package main

import (
    "fmt"

    "math"
)

type Number interface {
    int64 | float64
}

func main() {
    // Initialize a map for the integer values
    ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    // Initialize a map for the float values
    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

    GenericsOrIntFloat(ints, floats)
    GenericsComparable(ints, floats)
    GenericsInferred(ints, floats)
    noGenericsFuncs(ints, floats)
    noGenericsInterface(ints, floats)
}

func GenericsOrIntFloat(ints map[string]int64, floats map[string]float64) {
    fmt.Sprintf("Generic Sums: %v and %v\n",
        SumIntsOrFloats[string, int64](ints),
        SumIntsOrFloats[string, float64](floats))
}

func GenericsComparable(ints map[string]int64, floats map[string]float64) {
    fmt.Sprintf("Generic Sums with Constraint: %v and %v\n",
        SumNumbers(ints),
        SumNumbers(floats))
}

func GenericsInferred(ints map[string]int64, floats map[string]float64) {
    fmt.Sprintf("Generic Sums, type parameters inferred: %v and %v\n",
        SumIntsOrFloats(ints),
        SumIntsOrFloats(floats))
}

func noGenericsFuncs(ints map[string]int64, floats map[string]float64) {
    fmt.Sprintf("Non-Generic Sums: %v and %v\n",
        SumInts(ints),
        SumFloats(floats))
}

func noGenericsInterface(ints map[string]int64, floats map[string]float64) {
    fmt.Sprintf("Non-Generic Interface Sums: %v and %v\n",
        Sum(ints),
        Sum(floats))
}

func Sum(m interface{}) (s int64) {
    switch m.(type) {
    case map[string]int:
        for _, v := range m.(map[string]int) {
            s += int64(v)
        }
    case map[string]int64:
        //println("map int64")
        for _, v := range m.(map[string]int64) {
            s += v
        }
    case map[string]float64:
        for _, v := range m.(map[string]float64) {
            s += int64(math.Round(v))
        }
    default:
        // println("no found")

    }
    return s
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}

// SumIntsOrFloats sums the values of map m. It supports both floats and integers
// as map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

// SumNumbers sums the values of map m. Its supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

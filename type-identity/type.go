// Go in action
// @jeffotoni
// 2019-01-16

package main

import "fmt"

type (
	A0 = []string
	A1 = A0
	A2 = struct{ a, b int }
	A3 = int
	A4 = func(A3, float64) *A0
	A5 = func(x int, _ float64) *[]string
)

type (
	B0 A0
	B1 []string
	B2 struct{ a, b int }
	B3 struct{ a, c int }
	B4 func(int, float64) *B0
	B5 func(x int, y float64) *A1
)

type C0 = B0

func main() {
	var str C0
	str = append(str, "@jeffotoni")
	fmt.Println("types", str)
}

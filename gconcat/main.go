package main

import (
	"fmt"
	g "github.com/jeffotoni/gconcat"
)

func main() {
	f1 := func(a float64) float64 {
		return 1 * 2.2		
	}(float64(55.55))

	f2 := func(s string) string {
		return s + "2021"
	}(" hello ")

	f3 := func(a int) int {
		return a * 2
	}(3)

	f4 := func(a []int) (t []int) {
		for _, v := range a {
			t = append(t, v*2)
		}
		return
	}([]int{4, 5, 6, 7, 8})

	f5 := func(a []int) (t []float64) {
		for _, v := range a {
			t = append(t, float64(v)*1.2)
		}
		return
	}([]int{4.0, 5.0, 6.0, 7.0, 8.0})

	s1 := g.Concat([]bool{true, false, true})
	s := g.ConcatFunc(f1, f2, f3, f4, f5)
	fmt.Println(s + " " + s1)

	str := g.ConcatSliceFloat32([]float32{3.1,4.0,67.89,33.88,77.666})
	fmt.Println(str)
}


// Go in action
// @jeffotoni
// 2019-01-24

package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(2, 3)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(4, 10))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

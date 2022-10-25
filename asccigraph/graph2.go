package main

import (
	"fmt"
	"math"

	"github.com/guptarohit/asciigraph"
)

func main() {
	var data []float64

	// sine curve
	for i := 0; i < 105; i++ {
		data = append(data, 15*math.Sin(float64(i)*((math.Pi*4)/120.0)))
	}
	graph := asciigraph.Plot(data, asciigraph.Height(30))

	fmt.Println(graph)
}

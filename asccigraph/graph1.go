package main

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

func main() {
	data := []float64{3, 4, 9, 6, 2, 4, 5, 8, 5, 10, 2, 7, 2, 5, 6, 7, 8, 9, 10, 11, 12, 4, 5, 6, 10, 1, 12}
	graph := asciigraph.Plot(data)

	fmt.Println(graph)
}

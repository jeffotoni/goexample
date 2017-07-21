/*
* Example tollbooth
*
* @package     main
* @author      @jeffotoni
* @size        16/07/2017
*
 */

package main

import "math"

type Adder interface {
	Add(a float64) float64
}

type floatAdder float64

func (f floatAdder) Add(a float64) float64 {

	return float64(f) + a
}

type intAdder int

func (i intAdder) Add(a float64) float64 {

	return float64(i) + a
}

func StdDev(a []Adder) float64 {

	var Prom float64
	sum := 0.0
	Total := 0.0
	n := len(a)
	N := float64(n)

	for i := 0; i < n; i++ {

		sum = a[i].Add(float64(i))
	}

	println("vet n: ", n)
	println("vet a: ", a)
	println("vet N: ", N)
	println("vet sum: ", sum)

	Prom = sum / N

	for i := 0; i < n; i++ {

		Total += a[i].Add(-Prom) * a[i].Add(-Prom)
	}

	Total = Total / N
	Total = math.Sqrt(Total)

	return Total
}

func main() {

	floats := []Adder{floatAdder(1.0), floatAdder(2.0), floatAdder(3.0), floatAdder(4.0), floatAdder(5.0)}
	println(StdDev(floats))

	ints := []Adder{intAdder(1), intAdder(2), intAdder(3)}
	println(StdDev(ints))
}

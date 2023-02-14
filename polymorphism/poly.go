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

type AddFace interface {
	Add(a float64) float64
}

type floatAddFace float64

func (f floatAddFace) Add(a float64) float64 {

	return float64(f) + a
}

type intAddFace int

func (i intAddFace) Add(a float64) float64 {

	println("IntFace: ", a)
	println("IntFace + a: ", float64(i)+a)
	return float64(i) + a
}

func StdDev(a []AddFace) float64 {

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

	println("vet Total: ", Total)

	Total = Total / N
	Total = math.Sqrt(Total)

	return Total
}

func main() {

	floats := []AddFace{floatAddFace(1.0), floatAddFace(2.0), floatAddFace(3.0), floatAddFace(4.0), floatAddFace(5.0)}
	println(StdDev(floats))

	ints := []AddFace{intAddFace(1), intAddFace(2), intAddFace(3)}
	println(StdDev(ints))
}

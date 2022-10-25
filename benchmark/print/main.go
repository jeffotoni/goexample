package main

import (
	"math"
	"math/rand"
)

type Donut struct {
	Radius     float32
	Thick      float32
	Toppings   []string
	GlutenFree bool
	Hole       bool
	Filling    string
}

type Circle struct {
	r float64
}

func (c *Circle) area1() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) area2() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) area3() *float64 {
	s := math.Pi * c.r * c.r
	return &s
}

const maxToppings = 3

var radiuses = []float32{5, 10, 15}
var thicks = []float32{2, 3, 4}
var toppings = []string{"Chocolate", "Nuts", "Sugar", "Caramel"}
var fillings = []string{"", "Mermelade", "Chocolate", "Cream"}

var rnd = rand.New(rand.NewSource(321))

func RandomDonut1() *Donut {
	d := Donut{
		Radius: radiuses[rnd.Intn(len(radiuses))],
		Thick:  thicks[rnd.Intn(len(radiuses))],
	}
	if rnd.Int()%2 == 0 {
		d.GlutenFree = true
	} else {
		d.GlutenFree = false
	}
	if rnd.Int()%2 == 0 {
		d.Hole = true
	} else {
		d.Hole = false
		d.Filling = fillings[rnd.Intn(len(fillings))]
	}
	numToppints := rnd.Intn(maxToppings)
	d.Toppings = make([]string, 0, maxToppings)
	for i := 0; i < numToppints; i++ {
		d.Toppings = append(d.Toppings,
			toppings[rnd.Intn(len(toppings))])
	}
	return &d
}

func RandomDonut2() Donut {
	d := Donut{
		Radius: radiuses[rnd.Intn(len(radiuses))],
		Thick:  thicks[rnd.Intn(len(radiuses))],
	}
	if rnd.Int()%2 == 0 {
		d.GlutenFree = true
	} else {
		d.GlutenFree = false
	}
	if rnd.Int()%2 == 0 {
		d.Hole = true
	} else {
		d.Hole = false
		d.Filling = fillings[rnd.Intn(len(fillings))]
	}
	numToppints := rnd.Intn(maxToppings)
	d.Toppings = make([]string, 0, maxToppings)
	for i := 0; i < numToppints; i++ {
		d.Toppings = append(d.Toppings,
			toppings[rnd.Intn(len(toppings))])
	}
	return d
}

func main() {
	var CC = Circle{r: 55.44}
	_ = CC.area3()

	RandomDonut1()
	RandomDonut2()
}

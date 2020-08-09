package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int { return len(a) }

func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	family := []Person{
		{"Jeffotoni", 44},
		{"Arthur", 3},
		{"Francisco", 7},
	}
	sort.Sort(ByAge(family))
	fmt.Println(family) // [{Arthur 3} {Francisco 7} {Jeffotoni 44}]
}

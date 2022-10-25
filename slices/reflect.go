/*
* Golang map generic
*
* @package     main
* @author      @jeffotoni
* @size        2018
 */

package main

import (
	"fmt"
	"reflect"
)

type Mult interface {

	//func(x int) int { return x * 2 }
	Multiplicar()
}

func (Soma *Mult) Multiplicar(x int) {

	//= func(x int) int { return x * 2 }
}

// vetor string
type V []int

func (v *V) Map(args interface{}) {

	val := reflect.ValueOf(args)
	fmt.Println(val.Kind())
	if val.Kind() == reflect.Array {
		fmt.Println("len = ", val.Len())
		for i := 0; i < val.Len(); i++ {
			e := val.Index(i)
			switch e.Kind() {
			case reflect.Int:
				fmt.Printf("%v, ", e.Int())
			case reflect.Float32:
				fallthrough
			case reflect.Float64:
				fmt.Printf("%v, ", e.Float())
			default:
				panic(fmt.Sprintf("invalid Kind: %v", e.Kind()))
			}
		}
		fmt.Println()
	}
}

func main() {

	int_ary := [4]int{1, 2, 3, 4}
	float32_ary := [4]float32{1.1, 2.2, 3.3, 4.4}
	float64_ary := []float64{1.1, 2.2, 3.3, 4.4}

	// func(x int) int { return x * 2 })
	//V{1, 2, 3, 4}.Map()
	Map(float32_ary)
	Map(float64_ary)
}

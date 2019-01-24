// Go in action
// @jeffotoni
// 2019-01-16

package main

import "fmt"

func main() {

	a := [2][2]int{
		{33, 51},
		{72, 92}, // This trailing comma is mandatory
	}

	fmt.Println(a)
	fmt.Println(a[0][0])
	fmt.Println(a[0][1])
	fmt.Println(a[1][0])
	fmt.Println(a[1][1])

	a3 := [...][3]int{
		{10, 11, 20},
		{12, 14, 22},
		{30, 15, 24},
		{32, 16, 26},
		{34, 17, 28},
		{36, 18, 30},
	}
	fmt.Println(a3)
}

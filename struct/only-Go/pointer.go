// Go in action
// @jeffotoni
// 2019-01-16

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	var v =  Vertex{1, 2}
	v.X = 2
	v.Y = 3
	//p := &v
	//p.X = 1e9

	fmt.Sprintf("%v %v", v, v)	
	//fmt.Println(v)
	//fmt.Println(p.Y)
}

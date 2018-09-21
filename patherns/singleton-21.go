package main

import "fmt"

type Kitchen struct {
	numOfLamps int
}

type House struct {
	Kitchen
	numOfLamps int
}

func main() {
	h := House{Kitchen{2}, 10} //kitchen has 2 lamps, and the House has a total of 10 lamps

	fmt.Println("House h has this many lamps:", h.numOfLamps) //this is ok - the outer House's numOfLamps hides the other one.  Output is 10.

	fmt.Println("The Kitchen in house h has this many lamps:", h.Kitchen.numOfLamps) //we can still reach the number of lamps in the kitchen by using the type name h.Kitchen
}

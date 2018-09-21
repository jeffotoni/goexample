package main

import "fmt"

type Kitchen struct {
	numOfLamps int
}

type Bedroom struct {
	numOfLamps int
}

type House struct {
	Kitchen
	Bedroom
}

func main() {

	h := House{Kitchen{2}, Bedroom{3}} //kitchen has 2 lamps, Bedroom has 3 lamps

	fmt.Println("Ambiguous number of lamps:", h.Bedroom.numOfLamps) //this is an error due to ambiguousness - is it Kitchen.numOfLamps or Bedroom.numOfLamps
}

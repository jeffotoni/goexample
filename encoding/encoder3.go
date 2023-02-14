package main

import (
	"encoding/json"
	"fmt"
)

type Box struct {
	Width  int
	Height int
	Color  string
	Open   bool
}

func main() {
	// Create an instance of the Box struct.
	box := Box{
		Width:  10,
		Height: 20,
		Color:  "blue",
		Open:   false,
	}
	// Create JSON from the instance data.
	// ... Ignore errors.
	b, _ := json.Marshal(box)
	// Convert bytes to string.
	s := string(b)
	fmt.Println(s)
}

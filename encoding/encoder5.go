package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Positions []int
}

func main() {
	// This JSON contains an int array.
	text := "{\"Positions\": [100, 200, 300, -1]}"
	// Get bytes.
	bytes := []byte(text)

	// Unmarshal JSON to Result struct.
	var result Result
	json.Unmarshal(bytes, &result)

	// Our int array is filled.
	fmt.Printf("Positions = %v", result.Positions)
	fmt.Println()
	// Print int array length.
	fmt.Printf("Length = %v", len(result.Positions))
	fmt.Println()
}

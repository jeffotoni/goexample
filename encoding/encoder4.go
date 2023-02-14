package main

import (
	"encoding/json"
	"fmt"
)

type Language struct {
	Id   int
	Name string
}

func main() {

	// String contains two JSON rows.
	text := "[{\"Id\": 100, \"Name\": \"Go\"}, {\"Id\": 200, \"Name\": \"Java\"}]"
	// Get byte slice from string.
	bytes := []byte(text)

	// Unmarshal string into structs.
	var languages []Language
	json.Unmarshal(bytes, &languages)

	// Loop over structs and display them.
	for l := range languages {
		fmt.Printf("Id = %v, Name = %v", languages[l].Id, languages[l].Name)
		fmt.Println()
	}
}

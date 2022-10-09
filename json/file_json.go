package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// This code only return a error bcz it's not in a real scenario, but to use in real case, you can use your struct

type Character struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetJsonFile(path string) error {

	fileJsonData, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	characters := make([]Character, 0)

	if err = json.Unmarshal(fileJsonData, &characters); err != nil {
		return err
	}

	fmt.Print("\nCharacters:\n")
	for c := 0; c < len(characters); c++ {
		fmt.Printf("\nname: %s | age: %d \n", characters[c].Name, characters[c].Age)
	}

	return nil
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type myStruct struct {
	Data string `json:"data"`
}

func (s *myStruct) MarshalJSON() ([]byte, error) {
	return []byte(`{"data":"charlie"}`), nil
}

func (s *myStruct) UnmarshalJSON(b []byte) error {
	// Insert the string directly into the Data member
	return json.Unmarshal(b, &s.Data)
}

func main() {
	// Create a struct with initial content "alpha"
	ms := myStruct{"alpha"}

	// Replace content with "bravo" using custom UnmarshalJSON() (SUCCESSFUL)
	if err := json.NewDecoder(bytes.NewBufferString(`"bravo"`)).Decode(&ms); err != nil {
		log.Fatal(err)
	}

	// Use custom MarshalJSON() to get "charlie" back (UNSUCCESSFUL)
	if err := json.NewEncoder(os.Stdout).Encode(ms); err != nil {
		log.Fatal(err)
	}

	// Trying another method (UNSUCCESSFUL)
	if ret, err := json.Marshal(ms); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(ret))
	}

	// Verify that the Marshaler interface is correctly implemented
	var marsh json.Marshaler
	marsh = &ms
	ret, _ := marsh.MarshalJSON()
	fmt.Println(string(ret)) // Prints "charlie"
}

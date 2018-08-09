package main

import (
	"encoding/json"
	"fmt"
)

type ESModel struct {
	AB string `json:"a_b"`
}

type APIModel struct {
	AB string `json:"aB"`
}

func main() {

	b := []byte(`{
            "a_b": "c"
    }`)

	// var x ESModel
	// json.Unmarshal(b, &x)
	// b, _ = json.MarshalIndent(APIModel(x), "", "  ")

	var x APIModel
	json.Unmarshal(b, &x)
	b, _ = json.MarshalIndent(ESModel(x), "", "  ")

	fmt.Println(string(b))
}

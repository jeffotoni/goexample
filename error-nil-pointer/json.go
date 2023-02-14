package main

import (
	"encoding/json"
	"fmt"
)

type Number struct {
	N int
}

type NilableNumber struct {
	N *int
}

func main() {
	zeroJSON := []byte(`{"N": 0}`)
	emptyJSON := []byte(`{}`)

	var zeroNumber Number
	json.Unmarshal(zeroJSON, &zeroNumber)
	var emptyNumber Number
	json.Unmarshal(emptyJSON, &emptyNumber)
	fmt.Println(zeroNumber.N, emptyNumber.N) // output: 0 0

	var zeroNilable NilableNumber
	json.Unmarshal(zeroJSON, &zeroNilable)
	var emptyNilable NilableNumber
	json.Unmarshal(emptyJSON, &emptyNilable)
	fmt.Println(*zeroNilable.N, emptyNilable.N) // output: 0
}

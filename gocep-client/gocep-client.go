package main

import (
	"fmt"
	"github.com/jeffotoni/gocep/pkg/cep"
)

func main() {

	result, err := cep.Search("06233903")
	fmt.Println(err)
	fmt.Println(result)
}

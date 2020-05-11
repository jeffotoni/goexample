package main

import (
	"fmt"
	i "interfaces"
)

func main() {
	fmt.Println(i.Create(i.UserSet("jeffotoni")))
	// endereco
	fmt.Println(i.Create(i.EndSet("Belo Horizonte")))
}

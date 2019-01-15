package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Lendo Ambiente:")
	fmt.Println(os.Getenv("AWS_REGION"))
	fmt.Println(os.Getenv("AWS_TEST"))
}

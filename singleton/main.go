package main

import "fmt"
import "pkg/singleton"

func main() {
	fmt.Println("Testando singleton!")
	singleton.GetInstance()
}

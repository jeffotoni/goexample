package main

import (
	"fmt"

	"bitbucket.org/bertimus9/systemstat"
)

var sample systemstat.MemSample

// This example shows how easy it is to get memory information
func main() {
	sample = systemstat.GetMemSample()
	fmt.Println("Total available RAM in kb:", sample.MemTotal, "k total")
	fmt.Println("Used RAM in kb:", sample.MemUsed, "k used")
	fmt.Println("Free RAM in kb:", sample.MemFree, "k free")
	fmt.Printf("The output is similar to, but somewhat different than:\n\ttop -n1 | grep Mem:\n")
}

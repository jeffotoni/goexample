package main

import (
	"fmt"
	"runtime"
)

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

func main() {

	runtime.GOMAXPROCS(-1)

	fmt.Printf("GOMAXPROCS is %d\n", getGOMAXPROCS())
	fmt.Println("CPUS: ", runtime.NumCPU())
}

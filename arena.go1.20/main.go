package main

// GOEXPERIMENT=arenas go run main.go

import "arena"
import "fmt"

type T struct{
	Foo string
	Bar [16]byte
}

func main() {
	mem := arena.NewArena()
	obj1 := arena.New[T](mem) // arena-allocated
	obj2 := arena.Clone(obj1) // heap-allocated
	fmt.Println("obj1:", obj1)
	fmt.Println("obj2:", obj2)

	fmt.Println(obj2 == obj1) // false
	mem.Free()

	fmt.Println("obj1:", obj1)
	fmt.Println("obj2:", obj2)

}

/*func processRequest(req *http.Request) {
	// Create an arena in the beginning of the function.
	mem := arena.NewArena()
	// Free the arena in the end.
	defer mem.Free()

	// Allocate a bunch of objects from the arena.
	for i := 0; i < 10; i++ {
		obj := arena.New[T](mem)
	}

	// Or a slice with length and capacity.
	slice := arena.MakeSlice[T](mem, 100, 200)
}*/




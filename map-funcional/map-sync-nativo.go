package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("sync.Map test (Go 1.9+ only)")

	// Create
	// the threadsafe map.
	var sm sync.Map

	// uma alternativa
	// para fazer a
	// chamada
	//sm := new(SyncMap)
	//sm.stringernal = make(map[string]string)

	key := "login"
	value := "jeffotoni"

	sm.Store(key, value)

	// buscando o conteudo
	result, ok := sm.Load(key)

	if ok {
		fmt.Println("Value Load:", result)
	} else {
		fmt.Println("Nao encontrou a key: ", key)
	}
}

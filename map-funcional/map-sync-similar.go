package main

import (
	"fmt"
	"sync"
)

type SyncMap struct {
	mutex sync.Mutex
	sync.RWMutex
	stringernal map[string]string
}

//var sm SyncMap

func NewSyncMap() *SyncMap {
	return &SyncMap{
		stringernal: make(map[string]string),
	}
}

func (rm *SyncMap) Load(key string) (value string, ok bool) {

	fmt.Println("store load: ", rm.stringernal)

	rm.RLock()
	result, ok := rm.stringernal[key]
	rm.RUnlock()

	return result, ok
}

func (rm *SyncMap) Delete(key string) {
	rm.Lock()
	delete(rm.stringernal, key)
	rm.Unlock()
}

func (rm *SyncMap) Store(key, value string) {

	rm.Lock()
	rm.stringernal[key] = value

	// show no console
	fmt.Println("added value: ", value, "for key: ", key)

	rm.Unlock()
}

func main() {

	// uma forma mais
	// elegante
	// de fazer
	// a chamada
	sm := NewSyncMap()

	// uma alternativa
	// para fazer a
	// chamada
	//sm := new(SyncMap)
	//sm.stringernal = make(map[string]string)

	key := "login"
	value := "jeffotoni"

	fmt.Println(":::::Similar sync.Map:::::")

	sm.Store(key, value)

	// buscando o conteudo
	result, ok := sm.Load(key)

	if ok {
		fmt.Println("Value Load:", result)
	} else {
		fmt.Println("Nao encontrou a key: ", key)
	}
}

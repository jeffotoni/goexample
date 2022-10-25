package main

import (
	"fmt"
	"sync"
)

func main() {
	regularMapUsage()
	syncMapUsage()
}

type RegularStringMap struct {
	sync.RWMutex
	internal map[string]string
}

func NewRegularStringMap() *RegularStringMap {
	return &RegularStringMap{
		internal: make(map[string]string),
	}
}

func (rm *RegularStringMap) Load(key string) (value string, ok bool) {
	rm.RLock()
	result, ok := rm.internal[key]
	rm.RUnlock()
	return result, ok
}

func (rm *RegularStringMap) Delete(key string) {
	rm.Lock()
	delete(rm.internal, key)
	rm.Unlock()
}

func (rm *RegularStringMap) Store(key, value string) {
	rm.Lock()
	rm.internal[key] = value
	rm.Unlock()
}

type RegularIntMap struct {
	sync.RWMutex
	internal map[int]int
}

func NewRegularIntMap() *RegularIntMap {
	return &RegularIntMap{
		internal: make(map[int]int),
	}
}

func (rm *RegularIntMap) Load(key int) (value int, ok bool) {
	rm.RLock()
	result, ok := rm.internal[key]
	rm.RUnlock()
	return result, ok
}

func (rm *RegularIntMap) Delete(key int) {
	rm.Lock()
	delete(rm.internal, key)
	rm.Unlock()
}

func (rm *RegularIntMap) Store(key, value int) {
	rm.Lock()
	rm.internal[key] = value
	rm.Unlock()
}

func regularMapUsage() {
	fmt.Println("Regular threatsafe map test")
	fmt.Println("---------------------------")

	// Create the threadsafe map.
	reg := NewRegularStringMap()

	// Fetch an item that doesn't exist yet.
	result, ok := reg.Load("hello")
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("value not found for key: `hello`")
	}

	// Store an item in the map.
	reg.Store("hello", "world")
	fmt.Println("added value: `world` for key: `hello`")

	// Fetch the item we just stored.
	result, ok = reg.Load("hello")
	if ok {
		fmt.Printf("result: `%s` found for key: `hello`\n", result)
	}

	fmt.Println("---------------------------")
	fmt.Println()
	fmt.Println()
}

func syncMapUsage() {

	fmt.Println("sync.Map test (Go 1.9+ only)")
	fmt.Println("----------------------------")

	// Create the threadsafe map.
	var sm sync.Map

	// Fetch an item that doesn't exist yet.
	result, ok := sm.Load("hello")
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("value not found for key: `hello`")
	}

	// Store an item in the map.
	sm.Store("hello", "world")
	fmt.Println("added value: `world` for key: `hello`")

	// Fetch the item we just stored.
	result, ok = sm.Load("hello")
	if ok {
		fmt.Printf("result: `%s` found for key: `hello`\n", result.(string))
	}

	fmt.Println("---------------------------")
}

package main

import (
	"fmt"
	"sync"
)

type MyStruct struct {
	ID   int    `json:"id"`
	User string `json:"user"`
	Data string `json:"data"`
}

type SyncMap struct {
	sync.RWMutex
	qws map[int]interface{}
}

func NewSyncMap() *SyncMap {
	return &SyncMap{
		qws: make(map[int]interface{}),
	}
}

func (rm *SyncMap) Load(key int) (qwsc interface{}, ok bool) {
	rm.RLock()
	qwsc, ok = rm.qws[key]
	rm.RUnlock()
	return qwsc, ok
}

func (rm *SyncMap) Delete(key int) {
	rm.Lock()
	delete(rm.qws, key)
	rm.Unlock()
}

func (rm *SyncMap) Store(key int, qwsc interface{}) {
	rm.Lock()
	rm.qws[key] = qwsc
	rm.Unlock()
}

var sm = NewSyncMap()

func main() {
	var my MyStruct
	my.User = "jeffotoni"
	my.Data = "my data here... All Things"
	sm.Store(1, my)

	q, ok := sm.Load(1)
	fmt.Println(ok)
	fmt.Println(q)

	// cast
	mys := q.(MyStruct)
	fmt.Println(mys.User)
	fmt.Println(mys.Data)
}

package syncmap

import (
	"sync"
)

type Growth struct {
	Country   string
	Indicator string
	Value     float32
	Year      int
}

type SyncMap struct {
	sync.Mutex
	App map[string]Growth
}

func NewSyncMap() *SyncMap {
	return &SyncMap{
		App: make(map[string]Growth),
	}
}

func (rm *SyncMap) Load(key string) (Appc Growth, ok bool) {
	rm.Lock()
	defer rm.Unlock()
	Appc, ok = rm.App[key]
	return
}

func (rm *SyncMap) Get(key string) Growth {
	var Appc Growth
	var ok bool
	rm.Lock()
	defer rm.Unlock()

	Appc, ok = rm.App[key]
	if ok {
		return Appc
	}
	return Growth{}
}

func (rm *SyncMap) Delete(key string) {
	rm.Lock()
	delete(rm.App, key)
	rm.Unlock()
}

func (rm *SyncMap) Store(key string, Appc Growth) {
	rm.Lock()
	rm.App[key] = Appc
	rm.Unlock()
}

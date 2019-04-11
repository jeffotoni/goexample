package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"

	"github.com/VictoriaMetrics/fastcache"
)

var cache = &fastcache.Cache{}
var cache2 = &fastcache.Cache{}

func main() {

	cache = fastcache.New(15e8)

	// A huge allocation to give the GC work to do
	//lotsOf := make([]*int, 15e8)
	//lotsOf := make([]int, 15e8)
	fmt.Println("Background GC work generated")

	// Force a GC to set a baseline we can see if we set GODEBUG=gctrace=1
	runtime.GC()

	// Use up all the CPU doing work that causes allocations that could be cleaned up by the GC.
	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work()
		}()
	}

	wg.Wait()

	// Make sure that this memory isn't optimised away
	//runtime.KeepAlive(lotsOf)
	runtime.KeepAlive(cache)
}

func work() {
	var i int
	for {

		cache2 = fastcache.New(15e6)
		//work := make([]*int, 1e6)
		//work := make([]int, 1e6)
		i++

		f := factorial(20)
		key := []byte(strconv.Itoa(i))
		value := []byte(strconv.Itoa(i))

		// set in memory
		cache.Set(key, value)

		if f != 2432902008176640000 {
			fmt.Println("fact: ", f)
			os.Exit(1)
		}

		// get in memory
		//fmt.Println("get:: ", string(cache.Get(nil, key)))
		//runtime.KeepAlive(work)
		runtime.KeepAlive(cache2)
	}
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

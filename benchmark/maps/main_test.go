package main

import (
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

// Used to prevent compiler optimizations to ensure no dead code elimination.
// These ensure our Load functions aren't eliminated because we capture the result.
var globalResult int
var globalResultChan = make(chan int, 100)

func nrand(n int) []int {
	i := make([]int, n)
	for ind := range i {
		i[ind] = rand.Int()
	}
	return i
}

func populateMap(n int, rm *RegularIntMap) []int {
	nums := nrand(n)
	for _, v := range nums {
		rm.Store(v, v)
	}
	return nums
}

func populateSyncMap(n int, sm *sync.Map) []int {
	nums := nrand(n)
	for _, v := range nums {
		sm.Store(v, v)
	}
	return nums
}

func BenchmarkStoreRegular(b *testing.B) {
	nums := nrand(b.N)
	rm := NewRegularIntMap()
	b.ResetTimer()
	for _, v := range nums {
		rm.Store(v, v)
	}
}

func BenchmarkStoreSync(b *testing.B) {
	nums := nrand(b.N)
	var sm sync.Map
	b.ResetTimer()
	for _, v := range nums {
		sm.Store(v, v)
	}
}

func BenchmarkDeleteRegular(b *testing.B) {
	nums := nrand(b.N)
	rm := NewRegularIntMap()
	for _, v := range nums {
		rm.Store(v, v)
	}

	b.ResetTimer()
	for _, v := range nums {
		rm.Delete(v)
	}
}

func BenchmarkDeleteSync(b *testing.B) {
	nums := nrand(b.N)
	var sm sync.Map
	for _, v := range nums {
		sm.Store(v, v)
	}

	b.ResetTimer()
	for _, v := range nums {
		sm.Delete(v)
	}
}

func BenchmarkLoadRegularFound(b *testing.B) {
	nums := nrand(b.N)
	rm := NewRegularIntMap()
	for _, v := range nums {
		rm.Store(v, v)
	}

	currentResult := 0
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		currentResult, _ = rm.Load(nums[i])
	}
	globalResult = currentResult
}

func BenchmarkLoadRegularNotFound(b *testing.B) {
	nums := nrand(b.N)
	rm := NewRegularIntMap()
	for _, v := range nums {
		rm.Store(v, v)
	}
	currentResult := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		currentResult, _ = rm.Load(i)
	}
	globalResult = currentResult
}

func BenchmarkLoadSyncFound(b *testing.B) {
	nums := nrand(b.N)
	var sm sync.Map
	for _, v := range nums {
		sm.Store(v, v)
	}
	currentResult := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, ok := sm.Load(nums[i])
		if ok {
			currentResult = r.(int)
		}
	}
	globalResult = currentResult
}

func BenchmarkLoadSyncNotFound(b *testing.B) {
	nums := nrand(b.N)
	var sm sync.Map
	for _, v := range nums {
		sm.Store(v, v)
	}
	currentResult := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, ok := sm.Load(i)
		if ok {
			currentResult = r.(int)
		}
	}
	globalResult = currentResult
}

func BenchmarkRegularStableKeys1(b *testing.B) {
	benchmarkRegularStableKeys(b, 1)
}

func BenchmarkRegularStableKeys2(b *testing.B) {
	benchmarkRegularStableKeys(b, 2)
}

func BenchmarkRegularStableKeys4(b *testing.B) {
	benchmarkRegularStableKeys(b, 4)
}

func BenchmarkRegularStableKeys8(b *testing.B) {
	benchmarkRegularStableKeys(b, 8)
}

func BenchmarkRegularStableKeys16(b *testing.B) {
	benchmarkRegularStableKeys(b, 16)
}

func BenchmarkRegularStableKeys32(b *testing.B) {
	benchmarkRegularStableKeys(b, 32)
}

func BenchmarkRegularStableKeys64(b *testing.B) {
	benchmarkRegularStableKeys(b, 64)
}

func benchmarkRegularStableKeys(b *testing.B, workerCount int) {
	runtime.GOMAXPROCS(workerCount)

	rm := NewRegularIntMap()
	populateMap(b.N, rm)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	// Holds our final results, to prevent compiler optimizations.
	globalResultChan = make(chan int, workerCount)

	b.ResetTimer()

	for wc := 0; wc < workerCount; wc++ {
		go func(n int) {
			currentResult := 0
			for i := 0; i < n; i++ {
				currentResult, _ = rm.Load(5)
			}
			globalResultChan <- currentResult
			wg.Done()
		}(b.N)
	}

	wg.Wait()
}

func BenchmarkSyncStableKeys1(b *testing.B) {
	benchmarkSyncStableKeys(b, 1)
}

func BenchmarkSyncStableKeys2(b *testing.B) {
	benchmarkSyncStableKeys(b, 2)
}

func BenchmarkSyncStableKeys4(b *testing.B) {
	benchmarkSyncStableKeys(b, 4)
}

func BenchmarkSyncStableKeys8(b *testing.B) {
	benchmarkSyncStableKeys(b, 8)
}

func BenchmarkSyncStableKeys16(b *testing.B) {
	benchmarkSyncStableKeys(b, 16)
}

func BenchmarkSyncStableKeys32(b *testing.B) {
	benchmarkSyncStableKeys(b, 32)
}

func BenchmarkSyncStableKeys64(b *testing.B) {
	benchmarkSyncStableKeys(b, 64)
}

func benchmarkSyncStableKeys(b *testing.B, workerCount int) {
	runtime.GOMAXPROCS(workerCount)

	var sm sync.Map
	populateSyncMap(b.N, &sm)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	// Holds our final results, to prevent compiler optimizations.
	globalResultChan = make(chan int, workerCount)

	b.ResetTimer()

	for wc := 0; wc < workerCount; wc++ {
		go func(n int) {
			currentResult := 0
			for i := 0; i < n; i++ {
				r, ok := sm.Load(5)
				if ok {
					currentResult = r.(int)
				}
			}
			globalResultChan <- currentResult
			wg.Done()
		}(b.N)
	}

	wg.Wait()
}

func BenchmarkRegularStableKeysFound1(b *testing.B) {
	benchmarkRegularStableKeysFound(b, 1)
}

func BenchmarkRegularStableKeysFound2(b *testing.B) {
	benchmarkRegularStableKeysFound(b, 2)
}

func BenchmarkRegularStableKeysFound4(b *testing.B) {
	benchmarkRegularStableKeysFound(b, 4)
}

func BenchmarkRegularStableKeysFound8(b *testing.B) {
	benchmarkRegularStableKeysFound(b, 8)
}

func BenchmarkRegularStableKeysFound16(b *testing.B) {
	benchmarkRegularStableKeysFound(b, 16)
}

func BenchmarkRegularStableKeysFound32(b *testing.B) {
	benchmarkRegularStableKeysFound(b, 32)
}

func BenchmarkRegularStableKeysFound64(b *testing.B) {
	benchmarkRegularStableKeysFound(b, 64)
}

func benchmarkRegularStableKeysFound(b *testing.B, workerCount int) {
	runtime.GOMAXPROCS(workerCount)

	rm := NewRegularIntMap()
	values := populateMap(b.N, rm)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	// Holds our final results, to prevent compiler optimizations.
	globalResultChan = make(chan int, workerCount)

	b.ResetTimer()

	for wc := 0; wc < workerCount; wc++ {
		go func(n int) {
			currentResult := 0
			for i := 0; i < n; i++ {
				currentResult, _ = rm.Load(values[i])
			}
			globalResultChan <- currentResult
			wg.Done()
		}(b.N)
	}

	wg.Wait()
}

func BenchmarkSyncStableKeysFound1(b *testing.B) {
	benchmarkSyncStableKeysFound(b, 1)
}

func BenchmarkSyncStableKeysFound2(b *testing.B) {
	benchmarkSyncStableKeysFound(b, 2)
}

func BenchmarkSyncStableKeysFound4(b *testing.B) {
	benchmarkSyncStableKeysFound(b, 4)
}

func BenchmarkSyncStableKeysFound8(b *testing.B) {
	benchmarkSyncStableKeysFound(b, 8)
}

func BenchmarkSyncStableKeysFound16(b *testing.B) {
	benchmarkSyncStableKeysFound(b, 16)
}

func BenchmarkSyncStableKeysFound32(b *testing.B) {
	benchmarkSyncStableKeysFound(b, 32)
}

func BenchmarkSyncStableKeysFound64(b *testing.B) {
	benchmarkSyncStableKeysFound(b, 64)
}

func benchmarkSyncStableKeysFound(b *testing.B, workerCount int) {
	runtime.GOMAXPROCS(workerCount)

	var sm sync.Map
	values := populateSyncMap(b.N, &sm)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	// Holds our final results, to prevent compiler optimizations.
	globalResultChan = make(chan int, workerCount)

	b.ResetTimer()

	for wc := 0; wc < workerCount; wc++ {
		go func(n int) {
			currentResult := 0
			for i := 0; i < n; i++ {
				r, ok := sm.Load(values[i])
				if ok {
					currentResult = r.(int)
				}
			}
			globalResultChan <- currentResult
			wg.Done()
		}(b.N)
	}

	wg.Wait()
}

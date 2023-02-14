package main

import (
	"runtime"
	"sync/atomic"
	"testing"
)

func SomaOne(nums ...interface{}) interface{} {
	var s uint64
	for _, n := range nums {
		s += uint64(n.(int))
	}
	return s
}

func SomaTotal(nums ...int) uint64 {
	var s uint64
	for _, n := range nums {
		s += uint64(n)
	}
	return s
}

func BenchmarkInterface(b *testing.B) {
	var s uint64
	for i := 0; i < b.N; i++ {
		s += SomaOne(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).(uint64)
		s += SomaOne(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).(uint64)
		s += SomaOne(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).(uint64)
		s += SomaOne(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).(uint64)
	}
}

func BenchmarkTypeFast(b *testing.B) {
	var s uint64
	for i := 0; i < b.N; i++ {
		s += SomaTotal(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		s += SomaTotal(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		s += SomaTotal(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		s += SomaTotal(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
}
func BenchmarkInterfaceGoroutines(b *testing.B) {
	b.SetParallelism(runtime.NumCPU())
	var s uint64
	ch, done := make(chan uint64), make(chan bool, 4)
	f := func() {
		for i := 0; i < b.N; i++ {
			ch <- SomaOne(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).(uint64)
		}
	}
	l := func() {
		for i := 0; i < b.N; i++ {
			s += <-ch
		}
		done <- true
	}

	for i := 0; i < 4; i++ {
		go f()
		go l()
	}

	for i := 0; i < 4; i++ {
		<-done
	}
}

func BenchmarkTypeGoroutines(b *testing.B) {
	b.SetParallelism(runtime.NumCPU())
	var s uint64
	ch, done := make(chan uint64), make(chan bool, 4)
	f := func() {
		for i := 0; i < b.N; i++ {
			ch <- SomaTotal(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		}
	}
	l := func() {
		for i := 0; i < b.N; i++ {
			s += <-ch
		}
		done <- true
	}

	for i := 0; i < 4; i++ {
		go f()
		go l()
	}

	for i := 0; i < 4; i++ {
		<-done
	}
}

func BenchmarkInterfaceGoroutinesBuffered(b *testing.B) {
	b.SetParallelism(runtime.NumCPU())
	var s uint64
	ch, done := make(chan uint64, b.N/4), make(chan bool, 4)
	f := func() {
		for i := 0; i < b.N; i++ {
			ch <- SomaOne(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).(uint64)
		}
	}
	l := func() {
		for i := 0; i < b.N; i++ {
			s += <-ch
		}
		done <- true
	}

	for i := 0; i < 4; i++ {
		go f()
		go l()
	}

	for i := 0; i < 4; i++ {
		<-done
	}
}

func BenchmarkTypeGoroutinesBuffered(b *testing.B) {
	b.SetParallelism(runtime.NumCPU())
	var s uint64
	ch, done := make(chan uint64, b.N/4), make(chan bool, 4)
	f := func() {
		for i := 0; i < b.N; i++ {
			ch <- SomaTotal(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		}
	}
	l := func() {
		for i := 0; i < b.N; i++ {
			s += <-ch
		}
		done <- true
	}

	for i := 0; i < 4; i++ {
		go f()
		go l()
	}

	for i := 0; i < 4; i++ {
		<-done
	}
}

func BenchmarkInterfaceGoroutinesAtomic(b *testing.B) {
	b.SetParallelism(runtime.NumCPU())
	var s uint64
	done := make(chan bool, 4)
	f := func() {
		for i := 0; i < b.N; i++ {
			atomic.AddUint64(&s, SomaOne(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).(uint64))
		}
		done <- true
	}

	for i := 0; i < 4; i++ {
		go f()
	}

	for i := 0; i < 4; i++ {
		<-done
	}
}

func BenchmarkTypeGoroutinesAtomic(b *testing.B) {
	b.SetParallelism(runtime.NumCPU())
	var s uint64
	done := make(chan bool, 4)
	f := func() {
		for i := 0; i < b.N; i++ {
			atomic.AddUint64(&s, SomaTotal(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
		}
		done <- true
	}

	for i := 0; i < 4; i++ {
		go f()
	}

	for i := 0; i < 4; i++ {
		<-done
	}
}

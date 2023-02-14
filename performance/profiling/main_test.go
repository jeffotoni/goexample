package main

import (
	"sync"
	"testing"
)

var pi []int

func printlner(i ...int) {
	pi = i
}

type mySliceType struct {
	valuesGuard *sync.Mutex
	values      []int
}

func (s mySliceType) Get(idx int) int {
	s.valuesGuard.Lock()
	defer s.valuesGuard.Unlock()

	checkBuffer(s.values, idx)

	return s.values[idx]
}

func (s mySliceType) GetCh(ch chan int, idx int) {
	s.valuesGuard.Lock()
	defer s.valuesGuard.Unlock()

	checkBuffer(s.values, idx)

	ch <- s.values[idx]
}

func newMySliceType(values []int) mySliceType {
	return mySliceType{
		valuesGuard: &sync.Mutex{},
		values:      values,
	}
}

func fillBuffer(slice []int) map[int]int {
	result := map[int]int{}
	for i := 0; i < 100; i++ {
		for j := 0; j < len(slice); j++ {
			result[i*len(slice)+j] = slice[j]
		}
	}

	return result
}

func checkBuffer(slice []int, idx int) {
	buffer := make(map[int]int, len(slice)*100)
	buffer = fillBuffer(slice)
	for i := range buffer {
		if i == idx {
			return
		}
	}
}

func slicerInBounds(slice mySliceType) {
	for i := 0; i < 8; i++ {
		a0 := slice.Get(i*8 + 0)
		a1 := slice.Get(i*8 + 1)
		a2 := slice.Get(i*8 + 2)
		a3 := slice.Get(i*8 + 3)
		a4 := slice.Get(i*8 + 4)
		a5 := slice.Get(i*8 + 5)
		a6 := slice.Get(i*8 + 6)
		a7 := slice.Get(i*8 + 7)

		printlner(a0, a1, a2, a3, a4, a5, a6, a7)
	}
}

func slicerInBoundsChannels(slice mySliceType) {
	ch := make(chan int, 8)
	for i := 0; i < 8; i++ {
		go slice.GetCh(ch, i*8+0)
		go slice.GetCh(ch, i*8+1)
		go slice.GetCh(ch, i*8+2)
		go slice.GetCh(ch, i*8+3)
		go slice.GetCh(ch, i*8+4)
		go slice.GetCh(ch, i*8+5)
		go slice.GetCh(ch, i*8+6)
		go slice.GetCh(ch, i*8+7)

		a0 := <-ch
		a1 := <-ch
		a2 := <-ch
		a3 := <-ch
		a4 := <-ch
		a5 := <-ch
		a6 := <-ch
		a7 := <-ch

		printlner(a0, a1, a2, a3, a4, a5, a6, a7)
	}
}

func BenchmarkInBounds(b *testing.B) {
	var mySlice []int
	for i := 0; i < 99; i++ {
		mySlice = append(mySlice, i)
	}
	ms := newMySliceType(mySlice)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slicerInBounds(ms)
	}
}

func BenchmarkInBoundsChannels(b *testing.B) {
	var mySlice []int
	for i := 0; i < 99; i++ {
		mySlice = append(mySlice, i)
	}
	ms := newMySliceType(mySlice)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slicerInBoundsChannels(ms)
	}
}

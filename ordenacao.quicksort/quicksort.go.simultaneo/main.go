package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}

	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}

	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func quickSort_go(a []int, lo, hi int, done chan struct{}, depth int) {
	if lo >= hi {
		done <- struct{}{}
		return
	}

	depth--
	p := partition(a, lo, hi)
	if depth > 0 {
		childDone := make(chan struct{}, 2)
		go quickSort_go(a, lo, p-1, childDone, depth)
		go quickSort_go(a, p+1, hi, childDone, depth)

		<-childDone
		<-childDone
	} else {
		quickSort(a, lo, p-1)
		quickSort(a, p+1, hi)
	}

	done <- struct{}{}
}

func main() {
	qtd, _ := strconv.Atoi(os.Args[1])
	if qtd <= 1 {
		println("The value must be greater than 1!")
		return
	}

	start := time.Now()
	slice := generateSlice2(qtd)
	fmt.Println("slice load:", time.Now().Sub(start))

	done := make(chan struct{})
	start = time.Now()
	go quickSort_go(slice, 0, len(slice)-1, done, 5)
	<-done
	fmt.Println("Quicksort using goroutine:", time.Now().Sub(start))
}

func generateSlice2(qtd int) []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, 0, qtd)
	times := qtd
	for i := 0; i < times; i++ {
		val := rand.Intn(20000000)
		slice = append(slice, val)
	}
	return slice
}

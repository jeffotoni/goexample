package main

import "fmt"

func main() {
	x := FixBirdCountLog([]int{2, 3, 4, 5, 6, 7, 8})
	fmt.Println(x)
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
		if i&1 == 0 {
			//fmt.Println("i:", i, " i&1", i&1)
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}

func FixBirdCountLog2(birdsPerDay []int) []int {
	for i := range birdsPerDay {
		if i % 2 == 0 {
			//fmt.Println("i:", i, " i&1", i&1)
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
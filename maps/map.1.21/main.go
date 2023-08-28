package main

import (
	"fmt"
	"slices"
)

func main() {

	binarySrcTestInt := []int{3, 2, 4, 7, 3, 1, 2, 4, 6}
	i, found := slices.BinarySearch(binarySrcTestInt, 7)
	fmt.Printf("BinarySearch []int - found: %t | position: %d\n", found, i)

	// the slice must be ordered
	slices.Sort(binarySrcTestInt)
	fmt.Printf("%v\n", binarySrcTestInt)

	i, found = slices.BinarySearch(binarySrcTestInt, 7)
	fmt.Printf("BinarySearch - found: %t | position: %d\n", found, i)

	binarySrcTestStr := []string{"3", "2", "4", "7", "3", "1", "2", "4", "6"}
	i, found = slices.BinarySearch(binarySrcTestStr, "4")
	fmt.Printf("BinarySearch []string - found: %t | position: %d\n", found, i)

	newSliceInt := slices.Delete(binarySrcTestInt, 4, 5)
	fmt.Printf("New []int:    %v\n", binarySrcTestInt)
	fmt.Printf("Delete []int: %v\n", newSliceInt)

	newSliceReplace := slices.Replace(binarySrcTestInt, 1, 2, 8, 9)
	fmt.Printf("New []int:    %v\n", binarySrcTestInt)
	fmt.Printf("Replace []int: %v\n", newSliceReplace)

}

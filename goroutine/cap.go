package main

import "fmt"

func main() {
	sliceEx := make([]string, 0, 30)
	fmt.Printf("\nSlice: %d", cap(sliceEx))

	channelEx := make(chan string, 15)
	fmt.Printf("\nChannel: %d", cap(channelEx))

	var pointerEx *[40]string
	fmt.Printf("\nPointer: %d", cap(pointerEx))
}

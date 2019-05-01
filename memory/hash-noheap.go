package main

import (
	"fmt"

	"github.com/segmentio/fasthash/fnv1a"
)

func hashIt(in string) uint64 {
	out := fnv1a.HashString64(in)
	return out
}

func main() {
	s := "hello"
	fmt.Printf("The FNV64a hash of '%v' is '%v'\n", s, hashIt(s))
}

package main

import (
	"fmt"
	"hash/fnv"
)

func hashIt(in string) uint64 {

	h := fnv.New64a()
	h.Write([]byte(in))
	out := h.Sum64()
	return out
}

func main() {
	s := "hello"
	//writer := bufio.NewWriter(os.Stdout)
	//writer.WriteString(s)
	//defer writer.Flush()
	fmt.Printf("The FNV64a hash of '%v' is '%v'\n", s, hashIt(s))
}

package main

import (
	"bytes"
	"fmt"
	"sync"
)

var bufpool = sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 512)
		return &buf
	}}

func main() {
	bp := bufpool.Get().(*[]byte)
	b := *bp

	defer func() {
		*bp = b
		bufpool.Put(bp)
	}()

	buf := bytes.NewBuffer(b)
	fmt.Println(buf)
}

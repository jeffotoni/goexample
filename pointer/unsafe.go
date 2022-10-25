// Go Api server
// @jeffotoni
// 2019-03-10

package main

import (
	"bytes"
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := "hello jeff"
	b := []byte("hello jeff")
	p := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&a))))
	fmt.Println(bytes.Compare(b, p))
}

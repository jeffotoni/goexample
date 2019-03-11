// Go in action
// @jeffotoni
// 2019-03-10

package main

import (
	"fmt"
	"unsafe"
)

// ByteSliceToString is used when you really want to convert a slice // of bytes to a string without incurring overhead. It is only safe
// to use if you really know the byte slice is not going to change // in the lifetime of the string
func ByteSliceToString(bs []byte) string {
	// This is copied from runtime. It relies on the string
	// header being a prefix of the slice header!
	return *(*string)(unsafe.Pointer(&bs))
}

func main() {

	b := []byte("test parametr byte to string....")
	fmt.Println(*(*string)(unsafe.Pointer(&b)))
	// fmt.Println(ByteSliceToString(b))
}

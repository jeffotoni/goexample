// Go in Action
// @jeffotoni
// 2019-03-30

package main

import (
	"fmt"

	"github.com/shamaton/msgpack"
)

func main() {
	type Struct struct {
		Json string
	}
	v := Struct{Json: `{"key":"value","key2":"value2"}`}

	d, err := msgpack.Encode(v)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(d))

	r := Struct{}
	err = msgpack.Decode(d, &r)
	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}

// Go in Action
// @jeffotoni
// 2019-03-30

package main

import "fmt"
import "github.com/vmihailenco/msgpack"

func main() {
	type Item struct {
		Json string
	}

	b, err := msgpack.Marshal(&Item{Json: `{"key":"value","key2":"value2"}`})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	var item Item
	err = msgpack.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}

	fmt.Println(item)
	fmt.Println(item.Json)
	// Output: bar
}

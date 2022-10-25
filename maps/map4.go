// map with composite keys
package main

import (
	"fmt"
)

type key struct {
	key  string
	hash string
}

func main() {
	var m = map[key]string{
		{"foo", "1"}: "bar1",
		{"foo", "2"}: "bar2",
	}

	fmt.Println(m[key{"foo", "1"}])
	fmt.Println(m[key{"foo", "2"}])

	if _, ok := m[key{"foo", "3"}]; !ok {
		fmt.Printf("key not found")
	}
}

//Output:
/*
bar1
bar2
key not found
*/

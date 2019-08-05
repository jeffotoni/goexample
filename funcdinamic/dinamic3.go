package main

import (
	"fmt"
)

// Use map[string]interface{} to pair functions to name
// Could maybe use anonymous functions instead. Might be clean
// in certain cases
var funcMap = map[string]interface{}{
	"hello": hello,
	"name":  name,
}

func main() {
	callDynamically("hello")
	callDynamically("name", "Joe")
}

func callDynamically(name string, args ...interface{}) {
	switch name {
	case "hello":
		funcMap["hello"].(func())()
	case "name":
		funcMap["name"].(func(string))(args[0].(string))
	}

}

func hello() {
	fmt.Println("hello")
}

func name(name string) {
	fmt.Println(name)
}

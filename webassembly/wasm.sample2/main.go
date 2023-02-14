// @jeffotoni

package main

import (
	//"strconv"
	"syscall/js"
)

// set the value to js -> output
func setResult(val js.Value) {
	js.Global().Get("document").Call("getElementById", "result").Set("value", val)
}

func add(this js.Value, i []js.Value) interface{} {
	result := js.ValueOf(i[0].Int() + i[1].Int())
	setResult(result)
	return nil
}

func sub(this js.Value, i []js.Value) interface{} {
	result := js.ValueOf(i[0].Int() - i[1].Int())
	setResult(result)
	return nil
}

func div(this js.Value, i []js.Value) interface{} {
	result := js.ValueOf(i[0].Int() / i[1].Int())
	setResult(result)
	return nil
}

func mul(this js.Value, i []js.Value) interface{} {
	result := js.ValueOf(i[0].Int() * i[1].Int())
	setResult(result)
	return nil
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("sub", js.FuncOf(sub))
	js.Global().Set("div", js.FuncOf(div))
	js.Global().Set("mul", js.FuncOf(mul))
}

func main() {

	c := make(chan struct{}, 0)
	println("Go Webassembly inicialized!")
	registerCallbacks()

	<-c
}

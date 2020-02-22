package main

import (
	"fmt"
	"reflect"
	"runtime"
	"sync/atomic"
)

var finalizersCreated int64
var finalizersRan int64

func SetFinalizer(obj interface{}, finalizer interface{}) {
	finType := reflect.TypeOf(finalizer)
	funcType := reflect.FuncOf([]reflect.Type{finType.In(0)}, nil, false)
	f := reflect.MakeFunc(funcType, func(args []reflect.Value) []reflect.Value {
		finalizersRan++
		return reflect.ValueOf(finalizer).Call([]reflect.Value{args[0]})
	})
	runtime.SetFinalizer(obj, f.Interface())
	atomic.AddInt64(&finalizersCreated, 1)
}

func main() {
	v := "a"
	SetFinalizer(&v, func(a *string) {
		fmt.Println("Finalizer ran")
	})
	fmt.Println(finalizersRan, finalizersCreated)
	runtime.GC()
	fmt.Println(finalizersRan, finalizersCreated)
}

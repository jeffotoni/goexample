package main

import (
    "errors"
    "fmt"
    "reflect"
)

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {

    f := reflect.ValueOf(m[name])
    if len(params) != f.Type().NumIn() {
        err = errors.New("The number of params is not adapted.")
        return
    }

    in := make([]reflect.Value, len(params))
    for k, param := range params {
        in[k] = reflect.ValueOf(param)
    }

    result = f.Call(in)
    return
}

func main() {
    funcs := map[string]interface{}{"myfunc1": myfunc1, "myfunc2": myfunc2}
    Call(funcs, "myfunc1")
    Call(funcs, "myfunc2", 1, 2, 3)
}

func myfunc1() {
    // bla...bla...bla...
    fmt.Println("bla, bla.. myfunc1!!")
}
func myfunc2(a, b, c int) {
    // bla...bla...bla...
    fmt.Println("myfunc2:: ", a, b, c)
}

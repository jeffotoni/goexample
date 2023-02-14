// exemplo mapa generico
package main

import (
	"fmt"
	"reflect"
)

func main() {

	interf := []int{1, 2, 3, 4, 5, 6, 7, 8}

	r := genericMap(interf, func(x int) string {
		return fmt.Sprintf("%d", x)
	})

	fmt.Println(r)

	interfs := []string{"jefferson", "Otoni", "Lima", "é", "Um cara Legal"}

	rs := genericMap(interfs, func(x string) string {
		return "" + x
	})

	fmt.Println(rs)
}

func genericMap(arr interface{}, mapFunc interface{}) interface{} {

	funcValue := reflect.ValueOf(mapFunc)
	arrValue := reflect.ValueOf(arr)

	arrType := arrValue.Type()
	arrElemType := arrType.Elem()

	if arrType.Kind() != reflect.Array && arrType.Kind() != reflect.Slice {
		panic("O tipo de parâmetro não é nem array nem slice.")
	}

	funcType := funcValue.Type()

	// Verificando se o segundo argumento é função
	// E também verificando se sua assinatura é func ({type A}) {type B}
	if funcType.Kind() != reflect.Func || funcType.NumIn() != 1 || funcType.NumOut() != 1 {
		panic("Segundo argumento deve ser função de mapa")
	}

	if !arrElemType.ConvertibleTo(funcType.In(0)) {
		panic("O argumento da função de mapa não é compatível com o tipo do Array")
	}

	resultSliceType := reflect.SliceOf(funcType.Out(0))

	resultSlice := reflect.MakeSlice(resultSliceType, 0, arrValue.Len())

	for i := 0; i < arrValue.Len(); i++ {

		resultSlice = reflect.Append(resultSlice, funcValue.Call([]reflect.Value{arrValue.Index(i)})[0])
	}

	// Convering resulting slice back to generic interface.
	return resultSlice.Interface()
}

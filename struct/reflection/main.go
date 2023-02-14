// Name: jeffotoni
// Create struct in runtime
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type My struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	my := &My{Id: 87654, Name: "jeffotoni"}
	fmt.Printf("%+v\n", my)
	s := reflect.StructOf([]reflect.StructField{
		{
			Name: "Id",
			Type: reflect.TypeOf(int(0)),
			Tag:  `json:"id"`,
		},
		{
			Name: "Name",
			Type: reflect.TypeOf(""),
			Tag:  `json:"name"`,
		},
		{
			Name: "Cpf",
			Type: reflect.TypeOf(""),
			Tag:  `json:"cpf"`,
		},
	})

	d := reflect.New(s)
	d.Elem().Field(0).SetInt(int64(my.Id))
	d.Elem().Field(1).SetString(my.Name)
	d.Elem().Field(2).SetString("02997665478")
	myNew := d.Elem().Interface()
	fmt.Printf("%+v\n", d)

	b, err := json.Marshal(&myNew)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}

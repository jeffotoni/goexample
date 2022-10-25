package main

import (
	"reflect"
	"strings"
)

type MyStruct struct {
	Head string `json:"a1" xml:"x1"`
	Body string `json:"a2" xml:"x2"`
	Leg  string `json:"a3,omitempty" xml:"x3"`
}

func getField(tag, key string, s interface{}) (fieldname string) {
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		println("bad type")
		return
	}

	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		v := strings.Split(f.Tag.Get(key), ",")[0]
		if v == tag {
			println("field:", f.Name, " key:", key, " tag:", tag)
			return f.Name
		}
	}

	return ""
}

// OUT:
// field: Head  key: json  tag: a1
// field: Leg   key: json  tag: a3
// field: Body  key: json  tag: a2
// field: Body  key: xml   tag: x2

func main() {
	getField("a1", "json", MyStruct{})
	getField("a3", "json", MyStruct{})
	getField("a2", "json", MyStruct{})
	getField("x2", "xml", MyStruct{})
	getField("a99", "json", MyStruct{})
	// getField("a99", "json", "MyStruct{}")) // panic if not struct
}

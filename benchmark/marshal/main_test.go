package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"testing"
)

type Book struct {
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	Pages    int      `json:"num_pages"`
	Chapters []string `json:"chapters"`
}

type BookDef struct {
	Title    string   `msg:"title"`
	Author   string   `msg:"author"`
	Pages    int      `msg:"num_pages"`
	Chapters []string `msg:"chapters"`
}

/*
syntax = "proto2";
package main;

message BookProto {
  required string title = 1;
  required string author = 2;
  optional int64 pages = 3;
  repeated string chapters = 4;
}
*/

//$ sudo apt install protobuf-compiler
//$ go get -u github.com/golang/protobuf/protoc-gen-go
// protoc --go_out=. book.proto

// go get -u -t github.com/tinylib/msgp
/*
//go:generate msgp -tests=false
type BookDef struct {
	Title    string   `msg:"title"`
	Author   string   `msg:"author"`
	Pages    int      `msg:"num_pages"`
	Chapters []string `msg:"chapters"`
}
*/
// go generate

func generateObject() *Book {
	return &Book{
		Title:    "Computação quântica V.5",
		Author:   "Jefferson Otoni Lima",
		Pages:    1650,
		Chapters: []string{"Escala atômica,", "Arithmetic das partículas subatômicas"},
	}
}

func generateMessagePackObject() *BookDef {
	obj := generateObject()
	return &BookDef{
		Title:    obj.Title,
		Author:   obj.Author,
		Pages:    obj.Pages,
		Chapters: obj.Chapters,
	}
}

func generateProtoBufObject() *BookProto {
	obj := generateObject()
	return &BookProto{
		Title:    proto.String(obj.Title),
		Author:   proto.String(obj.Author),
		Pages:    proto.Int64(int64(obj.Pages)),
		Chapters: obj.Chapters,
	}
}

func BenchmarkJSONMarshal(b *testing.B) {
	obj := generateObject()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := json.Marshal(obj)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	out, err := json.Marshal(generateObject())
	if err != nil {
		panic(err)
	}

	obj := &Book{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = json.Unmarshal(out, obj)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkProtoBufMarshal(b *testing.B) {
	obj := generateProtoBufObject()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := proto.Marshal(obj)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkProtoBufUnmarshal(b *testing.B) {
	out, err := proto.Marshal(generateProtoBufObject())
	if err != nil {
		panic(err)
	}

	obj := &BookProto{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = proto.Unmarshal(out, obj)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMessagePackMarshal(b *testing.B) {
	obj := generateMessagePackObject()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := obj.MarshalMsg(nil)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMessagePackUnmarshal(b *testing.B) {
	obj := generateMessagePackObject()
	msg, err := obj.MarshalMsg(nil)
	if err != nil {
		panic(err)
	}

	obj = &BookDef{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err = obj.UnmarshalMsg(msg)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGobMarshal(b *testing.B) {
	obj := generateObject()

	enc := gob.NewEncoder(ioutil.Discard)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err := enc.Encode(obj)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGobUnmarshal(b *testing.B) {
	obj := generateObject()

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(obj)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		err = enc.Encode(obj)
		if err != nil {
			panic(err)
		}
	}

	dec := gob.NewDecoder(&buf)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = dec.Decode(&Book{})
		if err != nil {
			panic(err)
		}
	}
}

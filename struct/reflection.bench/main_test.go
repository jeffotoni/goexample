// @jeffotoni
// go test -run=. -bench=. -benchtime=3s -cou   nt 1 -benchmem
//
package test

import (
	"reflect"
	"testing"
)

type MyStruct struct {
	Name  string
	Age   int
	Code  string
	Point int
}

func BenchmarkReflect_New(b *testing.B) {
	var s *MyStruct
	sv := reflect.TypeOf(MyStruct{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sn := reflect.New(sv)
		s, _ = sn.Interface().(*MyStruct)
	}
	_ = s
}
func BenchmarkDirect_New(b *testing.B) {
	var s *MyStruct
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = new(MyStruct)
	}
	_ = s
}
func BenchmarkReflect_Set(b *testing.B) {
	var s *MyStruct
	sv := reflect.TypeOf(MyStruct{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sn := reflect.New(sv)
		s = sn.Interface().(*MyStruct)
		s.Name = "Jeff"
		s.Age = 18
		s.Code = "30067"
		s.Point = 100
	}
}
func BenchmarkReflect_SetFieldByName(b *testing.B) {
	sv := reflect.TypeOf(MyStruct{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sn := reflect.New(sv).Elem()
		sn.FieldByName("Name").SetString("Jeff")
		sn.FieldByName("Age").SetInt(18)
		sn.FieldByName("Code").SetString("30067")
		sn.FieldByName("Point").SetInt(100)
	}
}
func BenchmarkReflect_SetFieldByIndex(b *testing.B) {
	sv := reflect.TypeOf(MyStruct{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sn := reflect.New(sv).Elem()
		sn.Field(0).SetString("Jeff")
		sn.Field(1).SetInt(18)
		sn.Field(2).SetString("30067")
		sn.Field(3).SetInt(100)
	}
}
func BenchmarkDirect_Set(b *testing.B) {
	var s *MyStruct
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = new(MyStruct)
		s.Name = "Jeff"
		s.Age = 18
		s.Code = "30067"
		s.Point = 100
	}
}

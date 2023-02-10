package bench

import "testing"

func GenericCompare[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func DynCompare(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func TypedCompareStr(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func TypedCompareInt(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

var gr bool

var ga = []string{"foo", "bar", "baz"}
var gb = []string{"foo", "bar", "faz"}

var gan = []int{1, 2, 3}
var gbn = []int{1, 2, 4}

var gai = []interface{}{"foo", "bar", "baz"}
var gbi = []interface{}{"foo", "bar", "faz"}

var gain = []interface{}{1, 2, 3}
var gbin = []interface{}{1, 2, 4}

func BenchmarkTypedCompareStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gr = TypedCompareStr(ga, gb)
	}
}

func BenchmarkTypedCompareInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gr = TypedCompareInt(gan, gbn)
	}
}

func BenchmarkDynCompareStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gr = DynCompare(gai, gbi)
	}
}

func BenchmarkDynCompareInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gr = DynCompare(gain, gbin)
	}
}

func BenchmarkGenericCompareStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gr = GenericCompare(ga, gb)
	}
}

func BenchmarkGenericCompareInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gr = GenericCompare(gan, gbn)
	}
}

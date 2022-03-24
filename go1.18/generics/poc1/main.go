package main

type iMyinterface interface {
	int | int64 | int32
	~float64
	~string
	~[]byte
	MyMethod()
}

func NoGenericFuncInts(a, b []int) bool {
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

func NoGenericFuncStrs(a, b []string) bool {
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

func GenericsComparableSlice[T comparable](a, b []T) bool {
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

func GenericsAnySlice[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for range a {
		// list
	}
	return true
}

func NoGenericInterface(a, b interface{}) bool {
	switch a.(type) {
	case []int:
		aa := a.([]int)
		bb := b.([]int)
		if len(aa) != len(bb) {
			return false
		}

		for i, v := range a.([]int) {
			ii := b.([]int)
			if ii[i] != v {
				return false
			}
		}
		return true
	case []string:
		aa := a.([]string)
		bb := b.([]string)
		if len(aa) != len(bb) {
			return false
		}

		for i, v := range a.([]string) {
			ii := b.([]string)
			if ii[i] != v {
				return false
			}
		}
		return true
		return true
	}

	return false
}

var any1 string
var T string

func main() {
	println("version...")

	var t bool
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	b := a
	t = NoGenericFuncInts(a, b)
	println(t)

	s1 := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "jeff", "otoni", "go", "C"}
	s2 := s1
	t = NoGenericFuncStrs(s1, s2)
	println(t)

	t = GenericsComparableSlice(a, b)
	println(t)

	t = GenericsComparableSlice[string](s1, s2)
	println(t)

	t = NoGenericInterface(a, b)
	println(t)

	t = NoGenericInterface(s1, s2)
	println(t)

}

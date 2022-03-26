package main

import "testing"

var bt bool

func BenchmarkelBSlaice(b *testing.B) {
	//b.ResetTimer()
	var bb B
	for i := 0; i < b.N; i++ {
		bt = elBSlaice(bb)
	}
}

func BenchmarkelBAppend(b *testing.B) {
	//b.ResetTimer()
	var bb B
	for i := 0; i < b.N; i++ {
		bt = elBAppend(bb)
	}
}

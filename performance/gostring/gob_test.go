package gostring

import "testing"

var str string

const (
	concatSteps = 100
)

func TestGostring(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"gotest_1", args{1}},
		{"gotest_1", args{2}},
		{"gotest_1", args{3}},
		{"gotest_1", args{4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Gostring(tt.args.n)
		})
	}
}

func BenchmarkGostring(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < concatSteps; i++ {
			str = Gostring(i)
		}
	}
}

func BenchmarkGostring2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < concatSteps; i++ {
			str = Gostring2(i)
		}
	}
}

func BenchmarkGostring3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < concatSteps; i++ {
			str = Gostring3(i)
		}
	}
}

func BenchmarkGostring4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < concatSteps; i++ {
			str = Gostring4(i)
		}
	}
}

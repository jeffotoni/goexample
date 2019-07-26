package fib

import "testing"

// func TestFib(t *testing.T) {
// 	type args struct {
// 		n int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		// TODO: Add test cases.
// 		{"fib_0", args{0}, 0},
// 		{"fib_1", args{1}, 1},
// 		{"fib_2", args{2}, 1},
// 		{"fib_3", args{3}, 2},
// 		{"fib_4", args{4}, 3},
// 		{"fib_5", args{5}, 5},
// 		{"fib_6", args{6}, 10},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Fib(tt.args.n); got != tt.want {
// 				t.Errorf("Fib() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

var fibTests = []struct {
	n        int // input
	expected int // expected result
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
}

func TestFib(t *testing.T) {
	for _, tt := range fibTests {
		actual := Fib(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib1(b *testing.B) { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B) { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B) { benchmarkFib(3, b) }

// //func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
// //func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
// //func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }

var result int

func BenchmarkFibComplete(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = Fib(4)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

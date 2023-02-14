package main

import (
	"fmt"
	"testing"
)

func TestSum0(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"sum_1", args{0,1} , 1},
		{"sum_2", args{20,5} , 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			} else{
				t.Logf("PASS:%d=%d", got, tt.want)
			}
		})
	}
}


func TestSum1(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
	t.Run("a+b=20", func(t *testing.T) {
		s := Sum(10,10)
		if s != 20 {
			t.Errorf("Sum incorrect: got: %d, want: %d",s, 20)
		} else{
			t.Logf( "Sum(10,10) PASSED, expected %v and got value %v", 20, s )
		}
	})

}

func TestSum2(t *testing.T){
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 1},
		{1, 0, 1},
		{2, -2, 0},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			got := Sum(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
package main

import "testing"

func Test_worker(t *testing.T) {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	type args struct {
		id      int
		jobs    <-chan int
		results chan<- int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test_01",
			args: args{
				id:      1,
				jobs:    jobs,
				results: results,
			},
		},
	}
	for _, tt := range tests {
		t.Parallel()
		t.Run(tt.name, func(t *testing.T) {
			go worker(tt.args.id, tt.args.jobs, tt.args.results)
		})
	}
}

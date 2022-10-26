package main

import "testing"

func Test_tmpFunc(t *testing.T) {
	type args struct {
		apps []App
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "tmp_test_01",
			args: args{
				[]App{
					{
						Version: "1.0.0.0",
						Release: Release{
							Name: "One",
						},
						Title: "Lang C",
					},
				},
			},
		},

		{name: "tmp_test_02", args: args{[]App{
			{Version: "2.0.0.0", Release: Release{Name: "Two"},
				Title: "Lang Go",
			},
		}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmplFunc(tt.args.apps)
		})
	}
}

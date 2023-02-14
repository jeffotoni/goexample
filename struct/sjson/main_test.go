package main

import "testing"

func TestJsonM(t *testing.T) {
	type args struct {
		jstr string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test_struct_json_", args{jstr: `{"Plan":"planXy-939393","Loc":"","Discord":null}`}},
		{"test_struct_json_", args{jstr: `{"Plan":"","Loc":"BH","Discord":"here"}`}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JsonM(tt.args.jstr)
		})
	}
}

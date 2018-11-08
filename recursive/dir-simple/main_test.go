package main

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestDuracao(t *testing.T) {
	type args struct {
		seg string
	}
	tests := []struct {
		name         string
		args         args
		wantDuration time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDuration := Duracao(tt.args.seg); gotDuration != tt.wantDuration {
				t.Errorf("Duracao() = %v, want %v", gotDuration, tt.wantDuration)
			}
		})
	}
}

func TestLerFile(t *testing.T) {
	type args struct {
		Pathf string
	}
	tests := []struct {
		name string
		args args
		want []os.FileInfo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LerFile(tt.args.Pathf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LerFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLerDir(t *testing.T) {
	type args struct {
		Pathf string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LerDir(tt.args.Pathf)
		})
	}
}

// Golang In Action
// @package     main
// @author      @jeffotoni
// @size        2019

package main

import "testing"

func TestAnimal_Voar(t *testing.T) {
	tests := []struct {
		name string
		a    Animal
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Animal{}
			a.Voar()
		})
	}
}

func TestAnimal_Nadar(t *testing.T) {
	tests := []struct {
		name string
		a    Animal
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Animal{}
			a.Nadar()
		})
	}
}

func TestShowAnimal(t *testing.T) {
	type args struct {
		duck Duck
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShowAnimal(tt.args.duck)
		})
	}
}

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

package main

import (
	"testing"
)

func TestContainsNew(t *testing.T) {
	tests := []struct {
		algorithm string
		wantErr   bool
	}{
		{"AES", false},
		{"Blowfish", false},
		{"XYZ", true},
		{"DiffieHellman", false},
	}

	for _, tt := range tests {
		err := ContainsNew(tt.algorithm)
		if (err != nil) != tt.wantErr {
			t.Errorf("ContainsNew(%q) got error %v, want error: %v", tt.algorithm, err, tt.wantErr)
		}
	}
}

func TestContainsOld(t *testing.T) {
	tests := []struct {
		algorithm string
		wantErr   bool
	}{
		{"AES", false},
		{"Blowfish", false},
		{"XYZ", true},
		{"DiffieHellman", false},
	}

	for _, tt := range tests {
		err := ContainsOld(tt.algorithm)
		if (err != nil) != tt.wantErr {
			t.Errorf("ContainsOld(%q) got error %v, want error: %v", tt.algorithm, err, tt.wantErr)
		}
	}
}

func BenchmarkContainsNew(b *testing.B) {
	algorithm := "Blowfish"
	b.ResetTimer() // Resetamos o timer para excluir o tempo de inicialização
	for i := 0; i < b.N; i++ {
		ContainsNew(algorithm)
	}
}

func BenchmarkContainsOld(b *testing.B) {
	algorithm := "Blowfish"
	b.ResetTimer() // Resetamos o timer para excluir o tempo de inicialização
	for i := 0; i < b.N; i++ {
		ContainsOld(algorithm)
	}
}

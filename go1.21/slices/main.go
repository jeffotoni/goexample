package main

import (
	"fmt"
	"slices"
)

func main() {
	err := ContainsOld("NewBlowfish")
	fmt.Printf("%v\n", err)
	err = ContainsNew("Blowfish")
	fmt.Printf("%v\n", err)
}

// VERSAO NOVA
func ContainsNew(requestedAlgorithm string) error {
	supportedAlgorithms := []string{"AES", "Blowfish", "DES", "DSA", "DiffieHellman", "OAEP", "DESede"}
	if !slices.Contains(supportedAlgorithms, requestedAlgorithm) {
		return fmt.Errorf("unsupported algorithm: %s", requestedAlgorithm)
	}
	return nil
}

// VERSAO ANTIGA
func ContainsOld(requestedAlgorithm string) error {
	supportedAlgorithms := []string{"AES", "Blowfish", "DES", "DSA", "DiffieHellman", "OAEP", "DESede"}
	found := false
	for _, algo := range supportedAlgorithms {
		if algo == requestedAlgorithm {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("unsupported algorithm: %s", requestedAlgorithm)
	}
	return nil
}

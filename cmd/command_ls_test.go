package main

import "testing"

func TestRunLs(t *testing.T) {
	if err := runLs(); err != nil {
		t.Errorf("error: %v", err)
	}
}

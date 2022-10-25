package main

import "testing"

func TestLocked_add(t *testing.T) {
	c := locked{}
	c.add()
	if c.i != 1 {
		t.Errorf("Expected 1 but got %d", c.i)
	}
}

func BenchmarkLocked_add(t *testing.B) {
	c := locked{}
	for i := 0; i < t.N; i++ {
		c.add()
	}
	if c.i != t.N {
		t.Errorf("Expected %d, but got %d", t.N, c.i)
	}
}

func TestUnlocked_add(t *testing.T) {
	c := unlocked{}
	c.add()
	if c.i != 1 {
		t.Errorf("Expected 1 but got %d", c.i)
	}
}

func BenchmarkUnlocked_add(t *testing.B) {
	c := unlocked{}
	for i := 0; i < t.N; i++ {
		c.add()
	}
	if c.i != t.N {
		t.Errorf("Expected %d, but got %d", t.N, c.i)
	}
}

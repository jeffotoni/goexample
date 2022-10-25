package main

import "testing"

type fakeHTTPBin struct{}

func (l *fakeHTTPBin) Get(url string) ([]byte, error) {
	return []byte("Hello World test.."), nil
}

func TestBasics(t *testing.T) {
	expect := "Hello World test.."
	actual, _ := process(5, &fakeHTTPBin{})

	if actual != expect {
		t.Errorf("expected %s, actual %s", expect, actual)
	}
}

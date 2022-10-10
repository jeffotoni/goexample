package main

import "testing"

func TestGetJsonFile(t *testing.T) {

	err := GetJsonFile("./characters.json")

	if err != nil {
		t.Errorf("error: %v", err)
	}
}

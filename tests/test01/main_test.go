package main

import "testing"
import "strings"

func TestSum(t *testing.T) {

	a = App{}

	names := []struct {
		n string
	}{

		{"jeff"},
		{"jeff1"},
		{"jeff2"},
		{"jeff3"},
		{"jeff4"},
		{"jeff5"},
		{"jeff6"},
		{""},
		{" "},
		{"jeff8"},
	}

	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}

	for _, name := range names {

		a.Name = strings.Trim(name.n, "")

		if a.Name == "" {
			t.Errorf("Name (%s) incorreto", a.Name)
		}
	}
}

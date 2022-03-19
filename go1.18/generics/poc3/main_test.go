package main

import (
	"testing"
)

func TestStructGeneric(t *testing.T) {
	testcases := []struct {
		u MyStructGeneric[User]
		c MyStructGeneric[Company]
	}{
		{MyStructGeneric[User]{
			User{
				ID:"303939939393",
				Name: "jeffotoni",
				Cpf: 293399393,
			}},
			MyStructGeneric[Company]{
				Company{
					ID:123456,
					Phone: "5531234567897",
					CorporateName: "COMPANY LTDA.",
				},
			}},
	}
	for _, tc := range testcases {
		t.Log(tc.u.field)
		t.Log(tc.c.field)
	}
}

func FuzzStructGeneric(f *testing.F) {
	testcasesID := []string{"123456", "9499002", "093938"}
	for _, tc := range testcasesID {
		f.Add(tc)  // Use f.Add to provide a seed corpus
	}

	testcasesName := []string{"joao", "fabricio", "Neymar"}
	for _, tc := range testcasesName {
		f.Add(tc)  // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, orig string) {
		var u MyStructGeneric[User]
		u.field.ID = orig
		t.Log(u)
	})
}
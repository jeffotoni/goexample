package main

import "testing"

func TestLogin_SetLogin(t *testing.T) {
	type fields struct {
		Uuid   string
		User   string
		Status int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		// TODO: Add test cases.
		{"test_setlogin_1", fields{"xxxxxxxxxxx22333", "jefferson 1", 10}},
		{"test_setlogin_2", fields{"xxxxxxxxxxx344r7", "jefferson 2", 13}},
		{"test_setlogin_3", fields{"xxxxxxxxxxx666r3", "jefferson 3", 16}},
		{"test_setlogin_4", fields{"xxxxxxxxxxx66r33", "jefferson 4", 14}},
		{"test_setlogin_5", fields{"xxxxxxxxxxxt6653", "jefferson 5", 17}},
		{"test_setlogin_6", fields{"xxxxxxxxxxx83x34", "jefferson 6", 18}},
		{"test_setlogin_7", fields{"xxxxxxxxxxx2uy47", "jefferson 7", 19}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Login{
				Uuid:   tt.fields.Uuid,
				User:   tt.fields.User,
				Status: tt.fields.Status,
			}
			a.SetLogin()
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"jef"},
		{"lima"},
		{"go"},
		{"golang"},
		{"rust"},
		{"elixir"},
		{"Ocaml"},
		{"Scala"},
		{"Erlang"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

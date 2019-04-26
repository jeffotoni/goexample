package main

import "testing"

// go test -v
// go test -cpuprofile cpu.prof -memprofile mem.prof -bench pprof-2

// go test -cover -coverprofile=c.out
// go tool cover -html=c.out -o coverage.html

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

func Test_randomString1(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test_randomstring1_1", args{16}, "uuid1_1"},
		{"test_randomstring2_1", args{32}, "uuid1_2"},
		{"test_randomstring3_1", args{24}, "uuid1_3"},
		{"test_randomstring4_1", args{64}, "uuid1_4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomString1(tt.args.size); got != tt.want {
				t.Errorf("randomString1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randomString2(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test_randomstring2_1", args{16}, "uuid2_1"},
		{"test_randomstring2_2", args{32}, "uuid2_2"},
		{"test_randomstring3_1", args{64}, "uuid2_3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomString2(tt.args.l); got != tt.want {
				t.Errorf("randomString2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randInt(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randInt(tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("randInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyUuid(t *testing.T) {
	tests := []struct {
		name     string
		wantUuid string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUuid := MyUuid(); gotUuid != tt.wantUuid {
				t.Errorf("MyUuid() = %v, want %v", gotUuid, tt.wantUuid)
			}
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

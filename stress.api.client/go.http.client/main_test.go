package main

import (
	"net/http"
	"testing"

	"client2.http/internal/tests"
)

// go test -v -run ^TestHttpClient$
func TestHttpClient(t *testing.T) {
	type args struct {
		method      string
		ctype       string
		header      map[string]string
		url         string
		urlReq      string
		handlerfunc func(w http.ResponseWriter, r *http.Request)
	}
	tt := []struct {
		name     string
		args     args
		want     int //status code
		bodyShow bool
	}{
		{"test_get_", args{"GET", "application/json", nil, "/v1/client", "/v1/client", Get}, 200, false},
		{"test_get_", args{"GET", "application/json", nil, "/v1/client", "/v2/client", Get}, 200, false},
	}

	for _, tt := range tt {
		tt1 := tt
		t.Run(tt1.name, func(t *testing.T) {
			t.Parallel()
			tests.TestNewRequest(
				t,
				tt1.args.url,
				tt1.args.urlReq,
				tt1.args.method,
				tt1.args.handlerfunc,
				nil,
				tt1.args.ctype,
				tt1.args.header,
				tt1.want,
				tt1.bodyShow)
		})
	}
}

package tests

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRequest(
	t *testing.T,
	url,
	urlReq string,
	method string,
	handlerFunc func(w http.ResponseWriter, r *http.Request),
	data *bytes.Buffer,
	contentType string,
	header map[string]string,
	want int,
	bodyShow bool) {

	w := httptest.NewRecorder()
	var req *http.Request

	if data == nil {
		req = httptest.NewRequest(method, urlReq, nil)
	} else {
		req = httptest.NewRequest(method, urlReq, data)
	}

	if len(contentType) > 0 {
		req.Header.Set("Content-Type", contentType)
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	handlerFunc(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, want, resp.StatusCode)
	if bodyShow {
		var body []byte
		var err error
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Error ioutil.ReadAll:%s", err.Error())
			return
		}

		t.Log("http status:", resp.StatusCode)
		t.Log("\nResp :\n", string(body))
	}
}

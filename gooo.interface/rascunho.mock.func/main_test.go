package main

import (
	"net/http"
)

// MockClient is the mock client
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do is the mock client's `Do` func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	// coming soon!
	return req.Response, nil
}

//restclient.Client = &MockClient{}

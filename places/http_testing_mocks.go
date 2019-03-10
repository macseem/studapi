package places

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// MockResponseWriter is a mock for testing purpose
type MockResponseWriter struct {
	T             *testing.T
	expectedWrite []byte
}

// Header is a mock for testing purpose
func (m MockResponseWriter) Header() http.Header {
	h := make(map[string][]string)
	return h
}

// Write is mock for testing purpose
func (m MockResponseWriter) Write(data []byte) (int, error) {
	if string(m.expectedWrite) != string(data) {
		m.T.Error("Expected: " + string(m.expectedWrite) + " is not equal to: " + string(data))
	}
	return 0, nil
}

// WriteHeader is a mock for testing purpose
func (m MockResponseWriter) WriteHeader(statusCode int) {}

// MockClient is a mock of http.Client
type MockClient struct {
	T                *testing.T
	expectedResponse string
	expectedURL      string
}

// Get is a mock of http.Client.Get for testing purpose
// It uses expectedURL from MockClient to check the requestURL
// And returns a *http.Response with expectedResponse
func (m MockClient) Get(requestURL string) (*http.Response, error) {

	if requestURL != m.expectedURL {
		m.T.Error("Expected URL: " + requestURL + " is different")
	}
	bodyReadCloser := ioutil.NopCloser(strings.NewReader(m.expectedResponse))

	return &http.Response{Body: bodyReadCloser}, nil
}

package httpclient

import (
	"net/http"
	"testing"
)

//TestNewHTTPClient ...
func TestNewHTTPClient(t *testing.T) *HTTPClient {
	t.Helper()

	return &HTTPClient{
		baseURL:   "http://localhost",
		userAgent: "go client",
		client:    &http.Client{},
	}
}

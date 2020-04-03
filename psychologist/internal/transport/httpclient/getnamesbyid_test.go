package httpclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPClient_GetNamesByID(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		//assert.Equal(t, req.URL.String(), "/some/path")
		// Send response to be tested
		body, _ := json.Marshal(TestResponseGetNamesById(t))
		rw.Write(body)
	}))
	// Close the server when test finishes
	defer server.Close()

	hclient, err := New(server.URL, nil)
	assert.NoError(t, err)

	body, err := hclient.Do(TestSearchingClientsByID(t))
	assert.NoError(t, err)

	b, err := json.Marshal(TestResponseGetNamesById(t))
	assert.NoError(t, err)
	assert.Equal(t, body, b)
}

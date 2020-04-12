package httpclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

func TestHTTPClient_PsychologistName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/employees/75d2cdd6-cf69-44e7-9b28-c47792505d81/name")
		assert.Equal(t, req.Header.Get("Accept"), "application/json")
		assert.Equal(t, req.Header.Get("User-Agent"), "go client")
		assert.Equal(t, req.Header.Get("Content-Type"), "application/json")
		assert.Equal(t, req.Header.Get("X-Role"), "client")
		body := []byte(`{"id":"1","family_name":"Иванов","name":"Иван","patronomic":"Иванович"}`)
		rw.Write(body)
	}))

	defer server.Close()

	hclient, err := New(server.URL, "go client", server.Client())
	assert.NoError(t, err)

	body, err := hclient.PsychologistName("75d2cdd6-cf69-44e7-9b28-c47792505d81", rolePsychologist)
	assert.NoError(t, err)

	data, err := json.Marshal(&body)
	assert.NoError(t, err)

	wantedBody := []byte(`{"id":"1","family_name":"Иванов","name":"Иван","patronomic":"Иванович"}`)
	assert.Equal(t, data, wantedBody)
}

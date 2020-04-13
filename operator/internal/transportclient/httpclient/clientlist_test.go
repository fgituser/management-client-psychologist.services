package httpclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/fgituser/management-client-psychologist.services/operator/internal/model"
)

func TestHTTPClient_ClientsList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/client/list")
		assert.Equal(t, req.Header.Get("Accept"), "application/json")
		assert.Equal(t, req.Header.Get("User-Agent"), "go client")
		assert.Equal(t, req.Header.Get("Content-Type"), "application/json")
		assert.Equal(t, req.Header.Get("X-Role"), "admin")
		body := []byte(`[{"id":"48faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Гусев","name":"Евгений","patronomic":"Викторович","psychologist":{"id":"75d2cdd6-cf69-44e7-9b28-c47792505d81"}}]`)
		rw.Write(body)
	}))

	defer server.Close()

	hclient, err := New(server.URL, "go client", server.Client())
	assert.NoError(t, err)

	body, err := hclient.ClientsList()
	assert.NoError(t, err)

	data, err := json.Marshal(&body)
	assert.NoError(t, err)

	wantedBody := []byte(`[{"id":"48faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Гусев","name":"Евгений","patronomic":"Викторович","psychologist":{"id":"75d2cdd6-cf69-44e7-9b28-c47792505d81"}}]`)

	assert.Equal(t, data, wantedBody)
}

func TestHTTPClient_ClientsListByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/client/list_by_id")
		assert.Equal(t, req.Header.Get("Accept"), "application/json")
		assert.Equal(t, req.Header.Get("User-Agent"), "go client")
		assert.Equal(t, req.Header.Get("Content-Type"), "application/json")
		assert.Equal(t, req.Header.Get("X-Role"), "admin")
		body := []byte(`[{"id":"48faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Гусев","name":"Евгений","patronomic":"Викторович","psychologist":{"id":"75d2cdd6-cf69-44e7-9b28-c47792505d81"}}]`)
		rw.Write(body)
	}))

	defer server.Close()

	hclient, err := New(server.URL, "go client", server.Client())
	assert.NoError(t, err)

	body, err := hclient.ClientsListByID([]*model.Client{{ID: "48faa486-8e73-4c31-b10f-c7f24c115cda"}})
	assert.NoError(t, err)

	data, err := json.Marshal(&body)
	assert.NoError(t, err)

	wantedBody := []byte(`[{"id":"48faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Гусев","name":"Евгений","patronomic":"Викторович","psychologist":{"id":"75d2cdd6-cf69-44e7-9b28-c47792505d81"}}]`)

	assert.Equal(t, data, wantedBody)
}

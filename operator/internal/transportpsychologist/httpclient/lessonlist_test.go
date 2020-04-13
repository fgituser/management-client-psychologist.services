package httpclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

func TestHTTPClient_LessonList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/lessons/list")
		assert.Equal(t, req.Header.Get("Accept"), "application/json")
		assert.Equal(t, req.Header.Get("User-Agent"), "go client")
		assert.Equal(t, req.Header.Get("Content-Type"), "application/json")
		assert.Equal(t, req.Header.Get("X-Role"), "admin")
		body := []byte(`[{"client":{"id":"48faa486-8e73-4c31-b10f-c7f24c115cda"},"shedule":[{"employee":{"id":"50faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Гусев","name":"Евгений","patronomic":"Викторович"},"date_time":"2020-03-31T13:00:00Z"}]}]`)
		rw.Write(body)
	}))

	defer server.Close()

	hclient, err := New(server.URL, "go client", server.Client())
	assert.NoError(t, err)

	body, err := hclient.LessonList()
	assert.NoError(t, err)

	data, err := json.Marshal(&body)
	assert.NoError(t, err)

	wantedBody := []byte(`[{"client":{"id":"48faa486-8e73-4c31-b10f-c7f24c115cda"},"shedule":[{"psychologist":{"id":"50faa486-8e73-4c31-b10f-c7f24c115cda","family_name":"Гусев","name":"Евгений","patronomic":"Викторович"},"date_time":"2020-03-31T13:00:00Z"}]}]`)

	assert.Equal(t, data, wantedBody)
}

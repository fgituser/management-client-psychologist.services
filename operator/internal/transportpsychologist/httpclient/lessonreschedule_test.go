package httpclient

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alecthomas/assert"
)

func TestHTTPClient_LessonReschedule(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/employees/75d2cdd6-cf69-44e7-9b28-c47792505d81/clients/48faa486-8e73-4c31-b10f-c7f24c115cda/"+
			"lessons/datetime/2020-03-31%2013:00:00%20+0000%20UTC/reschedule/datetime/2020-03-31%2014:00:00%20+0000%20UTC/set")
		assert.Equal(t, req.Header.Get("Accept"), "application/json")
		assert.Equal(t, req.Header.Get("User-Agent"), "go client")
		assert.Equal(t, req.Header.Get("Content-Type"), "application/json")
		assert.Equal(t, req.Header.Get("X-Role"), "admin")
		rw.Write(nil)
	}))

	defer server.Close()

	hclient, err := New(server.URL, "go client", server.Client())
	assert.NoError(t, err)

	err = hclient.LessonReschedule("75d2cdd6-cf69-44e7-9b28-c47792505d81", "48faa486-8e73-4c31-b10f-c7f24c115cda",
		time.Date(2020, 3, 31, 13, 0, 0, 0, time.UTC),
		time.Date(2020, 3, 31, 14, 0, 0, 0, time.UTC))
	assert.NoError(t, err)
}

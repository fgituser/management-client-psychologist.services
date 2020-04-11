package httpclient

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		baseURL   string
		userAgent string
		client    *http.Client
	}
	tests := []struct {
		name    string
		args    args
		want    *HTTPClient
		wantErr bool
	}{
		{
			name:    "valid",
			args:    args{baseURL: "http://localhost", userAgent: "go client", client: &http.Client{}},
			want:    TestNewHTTPClient(t),
			wantErr: false,
		},
		{
			name:    "empty baseURL",
			args:    args{baseURL: "", userAgent: "go client", client: &http.Client{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "client is nil",
			args:    args{baseURL: "http://localhost", userAgent: "go client", client: nil},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.baseURL, tt.args.userAgent, tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPClient_Do(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/test_url")
		assert.Equal(t, req.Header.Get("Accept"), "application/json")
		assert.Equal(t, req.Header.Get("User-Agent"), "go client")
		assert.Equal(t, req.Header.Get("Content-Type"), "application/json")
		assert.Equal(t, req.Header.Get("X-Role"), "client")
		body := []byte("pong")
		rw.Write(body)
	}))

	defer server.Close()

	hclient, err := New(server.URL, "go client", server.Client())
	assert.NoError(t, err)

	body, err := hclient.Do(nil,
		"/test_url",
		"client")
	assert.NoError(t, err)

	wantedBody := []byte("pong")
	assert.NoError(t, err)
	assert.Equal(t, body, wantedBody)
}

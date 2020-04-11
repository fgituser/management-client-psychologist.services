package httpclient

import (
	"net/http"
	"reflect"
	"testing"
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

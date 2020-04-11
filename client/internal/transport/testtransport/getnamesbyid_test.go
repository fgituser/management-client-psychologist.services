package testtransport

import (
	"reflect"
	"testing"

	"github.com/fgituser/management-client-psychologist.services/client/internal/model"
)

func TestHTTPClient_GetNamesByID(t *testing.T) {
	type args struct {
		c          []*model.Client
		employeeID string
		userRole   string
	}
	tests := []struct {
		name    string
		h       *HTTPClient
		args    args
		want    []*model.Client
		wantErr bool
	}{
		{
			name: "valid",
			h: &HTTPClient{},
			args: args{
				c:          []*model.Client{},
				employeeID: "75d2cdd6-cf69-44e7-9b28-c47792505d81",
				userRole:   "test",
			},
			want:    TestClients(t),
			wantErr: false,
		},
		{
			name: "empty employyID",
			h: &HTTPClient{},
			args: args{
				c:          []*model.Client{},
				employeeID: "",
				userRole:   "test",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty userRole",
			h: &HTTPClient{},
			args: args{
				c:          []*model.Client{},
				employeeID: "75d2cdd6-cf69-44e7-9b28-c47792505d81",
				userRole:   "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "nil incomming clients",
			h: &HTTPClient{},
			args: args{
				c:          nil,
				employeeID: "75d2cdd6-cf69-44e7-9b28-c47792505d81",
				userRole:   "test",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HTTPClient{}
			got, err := h.GetNamesByID(tt.args.c, tt.args.employeeID, tt.args.userRole)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPClient.GetNamesByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPClient.GetNamesByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

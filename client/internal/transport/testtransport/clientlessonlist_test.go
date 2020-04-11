package testtransport

import (
	"reflect"
	"testing"
	"time"

	"github.com/fgituser/management-client-psychologist.services/client/internal/model"
)

func TestHTTPClient_ClientLessonList(t *testing.T) {
	type args struct {
		clientID       string
		psychologistID string
		userRole       string
	}
	tests := []struct {
		name    string
		h       *HTTPClient
		args    args
		want    []*model.Shedule
		wantErr bool
	}{
		{
			name: "valid",
			args: args{psychologistID: "80d2cdd6-cf69-44e7-9b28-c47792505d81", clientID: "75d2cdd6-cf69-44e7-9b28-c47792505d81", userRole: "client"},
			want: []*model.Shedule{
				{
					DateTime: time.Date(2020, 3, 31, 13, 0, 0, 0, time.UTC),
				},
			},
			wantErr: false,
		},
		{
			name:    "not valid empty psychologistID",
			args:    args{psychologistID: "", clientID: "75d2cdd6-cf69-44e7-9b28-c47792505d81", userRole: "client"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "not valid empty clientID",
			args:    args{psychologistID: "80d2cdd6-cf69-44e7-9b28-c47792505d81", clientID: "", userRole: "client"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HTTPClient{}
			got, err := h.ClientLessonList(tt.args.clientID, tt.args.psychologistID, tt.args.userRole)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPClient.ClientLessonList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPClient.ClientLessonList() = %v, want %v", got, tt.want)
			}
		})
	}
}

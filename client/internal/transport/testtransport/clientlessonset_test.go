package testtransport

import (
	"testing"
	"time"
)

func TestHTTPClient_ClientLessonSet(t *testing.T) {
	type args struct {
		clientID       string
		psychologistID string
		userRole       string
		dateTime       time.Time
	}
	tests := []struct {
		name    string
		h       *HTTPClient
		args    args
		wantErr bool
	}{
		{
			name:    "valid",
			h:       New(),
			args:    args{clientID: "75d2cdd6-cf69-44e7-9b28-c47792505d81", psychologistID: "58faa486-8e73-4c31-b10f-c7f24c115cda", userRole: "client", dateTime: time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)},
			wantErr: false,
		},
		{
			name:    "not valid clientDI",
			h:       New(),
			args:    args{clientID: " ", psychologistID: "58faa486-8e73-4c31-b10f-c7f24c115cda", userRole: "client", dateTime: time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)},
			wantErr: true,
		},
		{
			name:    "not valid psychologistID",
			h:       New(),
			args:    args{clientID: "75d2cdd6-cf69-44e7-9b28-c47792505d81", psychologistID: " ", userRole: "client", dateTime: time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HTTPClient{}
			if err := h.ClientLessonSet(tt.args.clientID, tt.args.psychologistID, tt.args.userRole, tt.args.dateTime); (err != nil) != tt.wantErr {
				t.Errorf("HTTPClient.ClientLessonSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

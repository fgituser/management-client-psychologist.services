package testtransport

import (
	"testing"
	"time"
)

func TestHTTPClient_LessonSet(t *testing.T) {
	type args struct {
		psychologistID string
		clientID       string
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
			args:    args{psychologistID: "80d2cdd6-cf69-44e7-9b28-c47792505d81", clientID: "75d2cdd6-cf69-44e7-9b28-c47792505d81", dateTime: time.Now()},
			wantErr: false,
		},
		{
			name:    "not valid psychologistID",
			h:       New(),
			args:    args{psychologistID: "", clientID: "75d2cdd6-cf69-44e7-9b28-c47792505d81", dateTime: time.Now()},
			wantErr: true,
		},
		{
			name:    "not valid clientID",
			h:       New(),
			args:    args{psychologistID: "80d2cdd6-cf69-44e7-9b28-c47792505d81", clientID: "", dateTime: time.Now()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HTTPClient{}
			if err := h.LessonSet(tt.args.psychologistID, tt.args.clientID, tt.args.dateTime); (err != nil) != tt.wantErr {
				t.Errorf("HTTPClient.LessonSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

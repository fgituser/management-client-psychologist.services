package testtransport

import (
	"testing"
	"time"
)

func TestHTTPClient_ClientLessonReschedule(t *testing.T) {
	type args struct {
		clientID       string
		psychologistID string
		userRole       string
		dateTimeOld    time.Time
		dateTimeNew    time.Time
	}
	tests := []struct {
		name    string
		h       *HTTPClient
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			h:    New(),
			args: args{clientID: "75d2cdd6-cf69-44e7-9b28-c47792505d81",
				psychologistID: "58faa486-8e73-4c31-b10f-c7f24c115cda",
				userRole:       "client",
				dateTimeOld:    time.Date(2020, 3, 31, 12, 0, 0, 0, time.UTC),
				dateTimeNew:    time.Date(2020, 3, 31, 13, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "not valid empty cleintID",
			h:    New(),
			args: args{clientID: "",
				psychologistID: "58faa486-8e73-4c31-b10f-c7f24c115cda",
				userRole:       "client",
				dateTimeOld:    time.Date(2020, 3, 31, 12, 0, 0, 0, time.UTC),
				dateTimeNew:    time.Date(2020, 3, 31, 13, 0, 0, 0, time.UTC),
			},
			wantErr: true,
		},
		{
			name: "not valid emtpy psychologistID",
			h:    New(),
			args: args{clientID: "75d2cdd6-cf69-44e7-9b28-c47792505d81",
				psychologistID: "",
				userRole:       "client",
				dateTimeOld:    time.Date(2020, 3, 31, 12, 0, 0, 0, time.UTC),
				dateTimeNew:    time.Date(2020, 3, 31, 13, 0, 0, 0, time.UTC),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HTTPClient{}
			if err := h.ClientLessonReschedule(tt.args.clientID, tt.args.psychologistID, tt.args.userRole, tt.args.dateTimeOld, tt.args.dateTimeNew); (err != nil) != tt.wantErr {
				t.Errorf("HTTPClient.ClientLessonReschedule() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

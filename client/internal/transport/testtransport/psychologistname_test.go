package testtransport

import (
	"reflect"
	"testing"

	"github.com/fgituser/management-client-psychologist.services/client/internal/model"
)

func TestHTTPClient_PsychologistName(t *testing.T) {
	type args struct {
		psychologistID string
		userRole       string
	}
	tests := []struct {
		name    string
		h       *HTTPClient
		args    args
		want    *model.Psychologist
		wantErr bool
	}{
		{
			name: "valid",
			h: New(),
			args: args{psychologistID: "75d2cdd6-cf69-44e7-9b28-c47792505d81", userRole: "client"},
			want: &model.Psychologist{
				ID:         "75d2cdd6-cf69-44e7-9b28-c47792505d81",
				FamilyName: "Васкецов",
				Name:       "Артем",
				Patronomic: "Викторович",
			},
			wantErr: false,
		},
		{
			name: "not valid",
			h: New(),
			args: args{psychologistID: "", userRole: "client"},
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HTTPClient{}
			got, err := h.PsychologistName(tt.args.psychologistID, tt.args.userRole)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPClient.PsychologistName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPClient.PsychologistName() = %v, want %v", got, tt.want)
			}
		})
	}
}

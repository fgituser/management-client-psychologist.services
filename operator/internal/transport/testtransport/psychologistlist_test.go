package testtransport

import (
	"reflect"
	"testing"

	"github.com/fgituser/management-client-psychologist.services/operator/internal/model"
)

func TestHTTPClient_PsychologistListByID(t *testing.T) {
	type args struct {
		psychologist []*model.Psychologist
	}
	tests := []struct {
		name    string
		h       *HTTPClient
		args    args
		want    []*model.Psychologist
		wantErr bool
	}{
		{
			name: "valid",
			h:    New(),
			args: args{psychologist: []*model.Psychologist{
				{
					ID: "60faa486-8e73-4c31-b10f-c7f24c115cda",
				},
			}},
			want: []*model.Psychologist{
				{
					ID:         "60faa486-8e73-4c31-b10f-c7f24c115cda",
					FamilyName: "Себастьянов",
					Name:       "Виктор",
					Patronomic: "Андреевич",
				},
			},
			wantErr: false,
		},
		{
			name: "not valid psychologist id",
			h:    New(),
			args: args{psychologist: []*model.Psychologist{
				{
					ID: "70faa486-8e73-4c31-b10f-c7f24c115cda",
				},
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "not valid psychologist nil",
			h:       New(),
			args:    args{psychologist: nil},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HTTPClient{}
			got, err := h.PsychologistListByID(tt.args.psychologist)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPClient.PsychologistListByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPClient.PsychologistListByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

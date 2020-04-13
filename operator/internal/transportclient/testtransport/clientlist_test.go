package testtransport

import (
	"reflect"
	"testing"

	"github.com/fgituser/management-client-psychologist.services/operator/internal/model"
)

func TestHTTPClient_ClientsList(t *testing.T) {
	tests := []struct {
		name    string
		h       *HTTPClient
		want    []*model.Client
		wantErr bool
	}{
		{
			name: "valid",
			h:    New(),
			want: []*model.Client{
				{
					ID:         "48faa486-8e73-4c31-b10f-c7f24c115cda",
					FamilyName: "Гусев",
					Name:       "Евгений",
					Patronomic: "Викторович",
					Psychologist: &model.Psychologist{
						ID: "60faa486-8e73-4c31-b10f-c7f24c115cda",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HTTPClient{}
			got, err := h.ClientsList()
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPClient.ClientsList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPClient.ClientsList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPClient_ClientsListByID(t *testing.T) {
	type args struct {
		clientsID []*model.Client
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
			h:    New(),
			args: args{clientsID: []*model.Client{{ID: "50faa486-8e73-4c31-b10f-c7f24c115cda"}}},
			want: []*model.Client{
				{
					ID:         "50faa486-8e73-4c31-b10f-c7f24c115cda",
					FamilyName: "Гусев",
					Name:       "Евгений",
					Patronomic: "Викторович",
					Psychologist: &model.Psychologist{
						ID: "60faa486-8e73-4c31-b10f-c7f24c115cda",
					},
				},
			},
			wantErr: false,
		},
		{
			name:    "not valid 1",
			h:       New(),
			args:    args{clientsID: []*model.Client{{ID: "58faa486-8e73-4c31-b10f-c7f24c115cda"}}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "not valid 2",
			h:       New(),
			args:    args{clientsID: nil},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HTTPClient{}
			got, err := h.ClientsListByID(tt.args.clientsID)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPClient.ClientsListByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPClient.ClientsListByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

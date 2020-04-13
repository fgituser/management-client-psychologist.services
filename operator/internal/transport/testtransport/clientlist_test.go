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

package testtransport

import (
	"reflect"
	"testing"
	"time"

	"github.com/fgituser/management-client-psychologist.services/operator/internal/model"
)

func TestHTTPClient_LessonList(t *testing.T) {
	tests := []struct {
		name    string
		h       *HTTPClient
		want    []*model.Employment
		wantErr bool
	}{
		{
			name: "valid",
			h:    New(),
			want: []*model.Employment{
				{
					Client: &model.Client{
						ID: "50faa486-8e73-4c31-b10f-c7f24c115cda",
					},
					Shedule: []*model.Shedule{
						{
							Psychologist: &model.Psychologist{
								ID:         "51faa486-8e73-4c31-b10f-c7f24c115cda",
								FamilyName: "Соболев",
								Name:       "Виктор",
								Patronomic: "Андреевич",
							},
							DateTime: time.Date(2020, 03, 31, 13, 0, 0, 0, time.UTC),
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HTTPClient{}
			got, err := h.LessonList()
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPClient.LessonList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPClient.LessonList() = %v, want %v", got, tt.want)
			}
		})
	}
}

package testtransport

import (
	"time"

	"github.com/fgituser/management-client-psychologist.services/operator/internal/model"
)

//LessonList ...
func (h *HTTPClient) LessonList() ([]*model.Employment, error) {
	return []*model.Employment{
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
	}, nil
}

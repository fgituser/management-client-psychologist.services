package testtransport

import (
	"github.com/fgituser/management-client-psychologist.services/operator/internal/model"
	"github.com/pkg/errors"
)

//PsychologistListByID get psychologist names from id
func (h *HTTPClient) PsychologistListByID(psychologist []*model.Psychologist) ([]*model.Psychologist, error) {
	if psychologist == nil || psychologist[0].ID != "60faa486-8e73-4c31-b10f-c7f24c115cda" {
		return nil, errors.New("[]*psychologits is nil or not valid psychologist.ID")
	}
	return []*model.Psychologist{
		{
			ID:         "60faa486-8e73-4c31-b10f-c7f24c115cda",
			FamilyName: "Себастьянов",
			Name:       "Виктор",
			Patronomic: "Андреевич",
		},
	}, nil
}

//PsychologistList get all psychologist
func (h *HTTPClient) PsychologistList() ([]*model.Psychologist, error) {
	return []*model.Psychologist{
		{
			ID:         "60faa486-8e73-4c31-b10f-c7f24c115cda",
			FamilyName: "Себастьянов",
			Name:       "Виктор",
			Patronomic: "Андреевич",
			Clients: []*model.Client{
				{
					ID: "50faa486-8e73-4c31-b10f-c7f24c115cda",
				},
			},
		},
	}, nil
}

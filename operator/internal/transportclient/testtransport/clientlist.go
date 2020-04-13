package testtransport

import (
	"errors"

	"github.com/fgituser/management-client-psychologist.services/operator/internal/model"
)

//ClientsList get clients
func (h *HTTPClient) ClientsList() ([]*model.Client, error) {
	return []*model.Client{
		{
			ID:         "48faa486-8e73-4c31-b10f-c7f24c115cda",
			FamilyName: "Гусев",
			Name:       "Евгений",
			Patronomic: "Викторович",
			Psychologist: &model.Psychologist{
				ID: "60faa486-8e73-4c31-b10f-c7f24c115cda",
			},
		},
	}, nil
}

//ClientsListByID get clients from id
func (h *HTTPClient) ClientsListByID(clientsID []*model.Client) ([]*model.Client, error) {
	if clientsID == nil || (clientsID != nil && clientsID[0].ID != "50faa486-8e73-4c31-b10f-c7f24c115cda") {
		return nil, errors.New("not valid clientsID")
	}
	return []*model.Client{
		{
			ID:         "50faa486-8e73-4c31-b10f-c7f24c115cda",
			FamilyName: "Гусев",
			Name:       "Евгений",
			Patronomic: "Викторович",
			Psychologist: &model.Psychologist{
				ID: "60faa486-8e73-4c31-b10f-c7f24c115cda",
			},
		},
	}, nil
}

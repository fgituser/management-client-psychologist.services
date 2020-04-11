package teststore

import (
	"strings"

	"github.com/fgituser/management-client-psychologist.services/client/internal/model"
	"github.com/pkg/errors"
)

//ClientsName get names clients from psychologistID
func (s *Store) ClientsName(psychologistID string) ([]*model.Client, error) {

	if strings.TrimSpace(psychologistID) != "75d2cdd6-cf69-44e7-9b28-c47792505d81a" {
		return nil, errors.New("psychologistID is empty")
	}

	return []*model.Client{
		{
			ID:         "48faa486-8e73-4c31-b10f-c7f24c115cda",
			FamilyName: "Гусев",
			Name:       "Евгений",
			Patronomic: "Викторович",
		},
	}, nil
}

//PsychologistID find psychologist attachment from client
func (s *Store) PsychologistID(clientID string) (string, error) {

	if strings.TrimSpace(clientID) != "48faa486-8e73-4c31-b10f-c7f24c115cda" {
		return "", errors.New("clientID not valid")
	}
	return "75d2cdd6-cf69-44e7-9b28-c47792505d81", nil
}

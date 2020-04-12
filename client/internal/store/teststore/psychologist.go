package teststore

import (
	"strings"

	"github.com/fgituser/management-client-psychologist.services/client/internal/model"
	"github.com/pkg/errors"
)

//ClientsName get names clients from psychologistID
func (s *Store) ClientsName(psychologistID string) ([]*model.Client, error) {

	if strings.TrimSpace(psychologistID) != "75d2cdd6-cf69-44e7-9b28-c47792505d81" {
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

	if strings.TrimSpace(clientID) != "75d2cdd6-cf69-44e7-9b28-c47792505d81" {
		return "", errors.New("clientID not valid")
	}
	return "58faa486-8e73-4c31-b10f-c7f24c115cda", nil
}

//IsAttachment check attachment client from psycholog
func (s *Store) IsAttachment(clientID, psychologistID string) (bool, error) {

	if strings.TrimSpace(clientID) != "48faa486-8e73-4c31-b10f-c7f24c115cda" {
		return false, errors.New("clientID is bad")
	}

	if strings.TrimSpace(psychologistID) != "75d2cdd6-cf69-44e7-9b28-c47792505d81" {
		return false, errors.New("psychologistID is bad")
	}
	return true, nil
}

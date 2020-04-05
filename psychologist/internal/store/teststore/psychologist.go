package teststore

import (
	"strings"
	"time"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/model"
	"github.com/pkg/errors"
)

//FindClients find all clients
func (s *Store) FindClients(employeeID string) ([]*model.Client, error) {
	if strings.TrimSpace(employeeID) == "" || employeeID != "75d2cdd6-cf69-44e7-9b28-c47792505d81" {
		return nil, errors.New("employeeID is empty or not valid")
	}
	return []*model.Client{
		{
			ID: "48faa486-8e73-4c31-b10f-c7f24c115cda",
		},
		{
			ID: "50faa486-8e73-4c31-b10f-c7f24c115cda",
		},
		{
			ID: "60faa486-8e73-4c31-b10f-c7f24c115cda",
		},
	}, nil
}

//LessonsList Get a list of your classes: date, client name
func (s *Store) LessonsList(employeeID string) ([]*model.Employment, error) {
	if strings.TrimSpace(employeeID) == "" || employeeID != "75d2cdd6-cf69-44e7-9b28-c47792505d81" {
		return nil, errors.New("employeeID is empty or not valid")
	}

	return []*model.Employment{
		{
			Client: &model.Client{
				ID: "48faa486-8e73-4c31-b10f-c7f24c115cda",
			},
			Shedule: []*model.Shedule{
				{
					DateTime: time.Date(2020, 3, 31, 13, 0, 0, 0, time.Local),
				},
			},
		},
	}, nil
}

package teststore

import (
	"strings"
	"time"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/model"
	"github.com/fgituser/management-client-psychologist.services/psychologist/pkg/datetime"
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

//SetLesson Schedule an activity with your client. Recording is possible at any time, including non-working
func (s *Store) SetLesson(employeeID, clientID string, dateTime time.Time) error {
	if strings.TrimSpace(employeeID) != "75d2cdd6-cf69-44e7-9b28-c47792505d81" ||
		strings.TrimSpace(clientID) != "48faa486-8e73-4c31-b10f-c7f24c115cda" {
		return errors.Errorf("an error accurred while set lesson, empty parametrs employeID:%v clientID:%v", employeeID, clientID)
	}

	_, _, err := datetime.DateTimeSplitUp(&dateTime)
	if err != nil {
		return errors.Wrap(err, "an error accurred while set lessons")
	}
	return nil
}

//CheckClientAttachment ...
func (s *Store) CheckClientAttachment(employeeID, clientID string) (bool, error) {
	if strings.TrimSpace(employeeID) != "75d2cdd6-cf69-44e7-9b28-c47792505d81" ||
		strings.TrimSpace(clientID) != "48faa486-8e73-4c31-b10f-c7f24c115cda" {
		return false, errors.Errorf("an error accured check client attachment: not valid parametrs employeeID:%v clientID:%v", employeeID, clientID)
	}

	return true, nil
}

//LessonIsBusy ...
func (s *Store) LessonIsBusy(employeeID string, dateTime time.Time) (bool, error) {
	if strings.TrimSpace(employeeID) == "" {
		return false, errors.Errorf("an error accured while check lesson free datetime: not valid parametrs employeeID:%v", employeeID)
	}
	//datetime busy
	return false, nil
}

//LessonCanceled canceled lesson
func (s *Store) LessonCanceled(employeeID string, dateTime time.Time) error {
	if strings.TrimSpace(employeeID) == "" {
		return errors.Errorf("an error accurred while caneled lesson, empty parametrs employeID:%v", employeeID)
	}
	_, _, err := datetime.DateTimeSplitUp(&dateTime)
	if err != nil {
		return errors.Wrap(err, "an error accurred while canceled lesson")
	}
	return nil
}

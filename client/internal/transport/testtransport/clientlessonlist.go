package testtransport

import (
	"errors"
	"strings"
	"time"

	"github.com/fgituser/management-client-psychologist.services/client/internal/model"
)

//ClientLessonList ...
func (h *HTTPClient) ClientLessonList(clientID, psychologistID, userRole string) ([]*model.Shedule, error) {
	if strings.TrimSpace(clientID) != "75d2cdd6-cf69-44e7-9b28-c47792505d81" || strings.TrimSpace(psychologistID) != "58faa486-8e73-4c31-b10f-c7f24c115cda" {
		return nil, errors.New("bad parametrs")
	}

	return []*model.Shedule{
		{
			DateTime: time.Date(2020, 3, 31, 13, 0, 0, 0, time.UTC),
		},
	}, nil

}

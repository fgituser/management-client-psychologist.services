package testtransport

import (
	"errors"
	"strings"
	"time"
)

//ClientLessonReschedule ...
func (h *HTTPClient) ClientLessonReschedule(clientID, psychologistID, userRole string, dateTimeOld, dateTimeNew time.Time) error {
	if strings.TrimSpace(psychologistID) != "58faa486-8e73-4c31-b10f-c7f24c115cda" ||
		strings.TrimSpace(clientID) != "75d2cdd6-cf69-44e7-9b28-c47792505d81" {
		return errors.New("an error occured set lesson")
	}
	return nil
}

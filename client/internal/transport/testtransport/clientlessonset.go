package testtransport

import (
	"errors"
	"strings"
	"time"
)

//ClientLessonSet ...
func (h *HTTPClient) ClientLessonSet(clientID, psychologistID, userRole string, dateTime time.Time) error {
	if strings.TrimSpace(psychologistID) != "80d2cdd6-cf69-44e7-9b28-c47792505d81" ||
		strings.TrimSpace(clientID) != "75d2cdd6-cf69-44e7-9b28-c47792505d81" {
		return errors.New("an error occured set lesson")
	}
	return nil
}

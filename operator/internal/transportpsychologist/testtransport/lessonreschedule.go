package testtransport

import (
	"time"

	"github.com/pkg/errors"
)

//LessonReschedule ...
func (h *HTTPClient) LessonReschedule(psychologistID, clientID string, dateTimeOld, dateTimeNew time.Time) error {
	if psychologistID != "80d2cdd6-cf69-44e7-9b28-c47792505d81" || clientID != "75d2cdd6-cf69-44e7-9b28-c47792505d81" {
		return errors.New("not valid incomming parametrs")
	}
	return nil
}

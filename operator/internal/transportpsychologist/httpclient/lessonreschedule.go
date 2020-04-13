package httpclient

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

//LessonReschedule ...
func (h *HTTPClient) LessonReschedule(psychologistID, clientID string, dateTimeOld, dateTimeNew time.Time) error {
	_, err := h.Do(nil, http.MethodPut,
		fmt.Sprintf("/employees/%v/clients/%v/lessons/datetime/%v/reschedule/datetime/%v/set", psychologistID, clientID, dateTimeOld, dateTimeNew), userRole)
	if err != nil {
		return errors.Wrap(err, "an error accured while reschedule lesson")
	}
	return nil
}

package httpclient

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

//ClientLessonReschedule ...
func (h *HTTPClient) ClientLessonReschedule(clientID, psychologistID, userRole string, dateTimeOld, dateTimeNew time.Time) error {
	if strings.TrimSpace(clientID) == "" || strings.TrimSpace(psychologistID) == "" {
		return errors.New("an error occured reschedule lesson")
	}

	_, err := h.Do(nil, http.MethodPut,
		fmt.Sprintf("/employees/%v/clients/%v/lessons/"+
			"datetime/%v/reschedule/datetime/%v/set", psychologistID, clientID, dateTimeOld, dateTimeNew),
		userRole)
	if err != nil {
		return errors.Wrap(err, "an error occured reschedule lesson")
	}
	return nil
}

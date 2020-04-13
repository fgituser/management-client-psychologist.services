package httpclient

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

//LessonSet ...
func (h *HTTPClient) LessonSet(psychologistID, clientID string, dateTime time.Time) error {
	_, err := h.Do(nil, http.MethodPost,
		fmt.Sprintf("/employees/%v/clients/%v/lessons/datetime/%v/set", psychologistID, clientID, dateTime), userRole)
	if err != nil {
		return errors.Wrap(err, "an error accured set lesson")
	}
	return nil
}

package httpclient

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

//ClientLessonSet ...
func (h *HTTPClient) ClientLessonSet(clientID, psychologistID, userRole string, dateTime time.Time) error {
	if strings.TrimSpace(clientID) == "" || strings.TrimSpace(psychologistID) == "" {
		return errors.New("an error occured set lesson")
	}

	_, err := h.Do(nil, http.MethodPost,
		fmt.Sprintf("/employees/%v/clients/%v/lessons/datetime/%v/set", psychologistID, clientID, dateTime),
		userRole)
	if err != nil {
		return errors.Wrap(err, "an error occured set lesson")
	}
	return nil
}

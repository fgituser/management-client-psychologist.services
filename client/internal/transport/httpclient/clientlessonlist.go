package httpclient

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/fgituser/management-client-psychologist.services/client/internal/model"
	"github.com/pkg/errors"
)

//TODO: check user role
const (
	rolePsychologist = "client"
)

type responseClientLessons struct {
	Datetime time.Time `json:"date_time"`
}

//ClientLessonList ...
func (h *HTTPClient) ClientLessonList(clientID, psychologistID, userRole string) ([]*model.Shedule, error) {
	if strings.TrimSpace(clientID) == "" || strings.TrimSpace(psychologistID) == "" {
		return nil, errors.New("an error accured get client lesson list")
	}

	res, err := h.Do(nil,
		fmt.Sprintf("/employees/%v/client/%v/lessons", psychologistID, clientID),
		userRole)
	if err != nil {
		return nil, errors.Wrap(err, "an error occured get client lesson list")
	}

	rClientLessons := make([]*responseClientLessons, 0)
	if err := json.Unmarshal(res, &rClientLessons); err != nil {
		return nil, errors.Wrap(err, "an error occured get client lesson list")
	}
	return convertClietnLessonsDtoToModelShedule(rClientLessons), nil

}

func convertClietnLessonsDtoToModelShedule(rs []*responseClientLessons) []*model.Shedule {
	schedule := make([]*model.Shedule, 0)
	for _, r := range rs {
		schedule = append(schedule, &model.Shedule{
			DateTime: r.Datetime,
		})
	}
	return schedule
}

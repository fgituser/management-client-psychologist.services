package httpclient

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fgituser/management-client-psychologist.services/operator/internal/model"
	"github.com/pkg/errors"
)

type respoonseLessonList struct {
	Client struct {
		ID string `json:"id"`
	} `json:"client"`
	Shedule []struct {
		Employee struct {
			ID         string `json:"id"`
			FamilyName string `json:"family_name"`
			Name       string `json:"name"`
			Patronomic string `json:"patronomic"`
		} `json:"employee"`
		DateTime time.Time `json:"date_time"`
	} `json:"shedule"`
}

//LessonList ...
func (h *HTTPClient) LessonList() ([]*model.Employment, error) {
	res, err := h.Do(nil, http.MethodGet, "/lessons/list", userRole)
	if err != nil {
		return nil, errors.Wrap(err, "an error accured while get all psychologist")
	}
	empl := make([]*respoonseLessonList, 0)
	if err := json.Unmarshal(res, &empl); err != nil {
		return nil, errors.Wrap(err, "an error accured while get all psychologist")
	}
	return respoonseLessonListToModelEmployment(empl), nil
}

func respoonseLessonListToModelEmployment(res []*respoonseLessonList) []*model.Employment {
	employment := make([]*model.Employment, 0)
	for _, r := range res {
		e := &model.Employment{
			Client: &model.Client{
				ID: r.Client.ID,
			},
		}
		for _, s := range r.Shedule {
			e.Shedule = append(e.Shedule, &model.Shedule{
				Psychologist: &model.Psychologist{
					ID:         s.Employee.ID,
					FamilyName: s.Employee.FamilyName,
					Name:       s.Employee.Name,
					Patronomic: s.Employee.Patronomic,
				},
				DateTime: s.DateTime,
			})
		}
		employment = append(employment, e)
	}
	return employment
}

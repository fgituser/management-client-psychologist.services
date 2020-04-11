package httpclient

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fgituser/management-client-psychologist.services/client/internal/model"
	"github.com/pkg/errors"
)

type responsePsychologistName struct {
	ID         string `json:"id,omitempty"`
	FamilyName string `json:"family_name,omitempty"`
	Name       string `json:"name,omitempty"`
	Patronomic string `json:"patronomic,omitempty"`
}

//PsychologistName ...
func (h *HTTPClient) PsychologistName(psychologistID, userRole string) (*model.Psychologist, error) {
	if strings.TrimSpace(psychologistID) == "" {
		return nil, errors.New("not valid psychologistID")
	}

	res, err := h.Do(nil,
		fmt.Sprintf("/employees/%v/name", psychologistID),
		userRole)
	if err != nil {
		return nil, errors.Wrap(err, "an error occured get psychologist")
	}

	var rPName = new(responsePsychologistName)
	if err := json.Unmarshal(res, &rPName); err != nil {
		return nil, errors.Wrap(err, "an error occured get get psychologist")
	}
	return convertPsychologistNameDtoToModelPsychologist(rPName), nil

}

func convertPsychologistNameDtoToModelPsychologist(rs *responsePsychologistName) *model.Psychologist {
	return &model.Psychologist{
		ID:         rs.ID,
		FamilyName: rs.FamilyName,
		Name:       rs.Name,
		Patronomic: rs.Patronomic,
	}
}

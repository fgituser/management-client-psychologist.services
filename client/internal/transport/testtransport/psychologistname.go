package testtransport

import (
	"strings"

	"github.com/fgituser/management-client-psychologist.services/client/internal/model"
	"github.com/pkg/errors"
)

//PsychologistName ...
func (h *HTTPClient) PsychologistName(psychologistID, userRole string) (*model.Psychologist, error) {
	if strings.TrimSpace(psychologistID) != "58faa486-8e73-4c31-b10f-c7f24c115cda" {
		return nil, errors.New("not valid psychologistID")
	}
	return &model.Psychologist{
		ID:         "58faa486-8e73-4c31-b10f-c7f24c115cda",
		FamilyName: "Васкецов",
		Name:       "Артем",
		Patronomic: "Викторович",
	}, nil

}

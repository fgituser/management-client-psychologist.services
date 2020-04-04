package testtransport

import (
	"errors"
	"strings"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/model"
)

//GetNamesByID getting client names by identifiers
func (h *HTTPClient) GetNamesByID(c []*model.Client, employeeID, userRole string) ([]*model.Client, error) {
	if c == nil || strings.TrimSpace(employeeID) == "" ||
		employeeID != "75d2cdd6-cf69-44e7-9b28-c47792505d81" ||
		strings.TrimSpace(userRole) == "" {
		return nil, errors.New("an error accured get names by id: bad parametrs")
	}


	return []*model.Client{
		{
			ID:         "48faa486-8e73-4c31-b10f-c7f24c115cda",
			FamilyName: "Гусев",
			Name:       "Евгений",
			Patronomic: "Викторович",
		},
		{
			ID:         "50faa486-8e73-4c31-b10f-c7f24c115cda",
			FamilyName: "Шмельцер",
			Name:       "Вячеслав",
			Patronomic: "Николаевич",
		},
		{
			ID:         "60faa486-8e73-4c31-b10f-c7f24c115cda",
			FamilyName: "Виевская",
			Name:       "Анастасия",
			Patronomic: "Федоровна",
		},
	}, nil
}

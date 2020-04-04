package teststore

import (
	"testing"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/model"
)

func TestClients(t *testing.T) []*model.Client {
	t.Helper()
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
	}
}

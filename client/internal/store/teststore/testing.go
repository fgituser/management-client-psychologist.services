package teststore

import (
	"testing"

	"github.com/fgituser/management-client-psychologist.services/client/internal/model"
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
	}
}

package httpclient

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/fgituser/management-client-psychologist.services/client/internal/model"
)

//TestClients test cliets
func TestSearchingClientsByID(t *testing.T) ([]byte, []*model.Client) {
	t.Helper()

	clientsID := []*model.Client{
		{
			ID: "48faa486-8e73-4c31-b10f-c7f24c115cda",
		},
		{
			ID: "50faa486-8e73-4c31-b10f-c7f24c115cda",
		},
		{
			ID: "60faa486-8e73-4c31-b10f-c7f24c115cda",
		},
	}
	data, err := json.Marshal(&clientsID)
	if err != nil {
		t.Fatal(err)
	}
	return data, clientsID
}

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

//TestResponseGetNamesById test response from GetNamesById
func TestResponseGetNamesById(t *testing.T) ([]byte, []*responseGetNamesByID) {
	t.Helper()
	clients := []*responseGetNamesByID{
		{
			ClinetID:   "48faa486-8e73-4c31-b10f-c7f24c115cda",
			FamilyName: "Гусев",
			Name:       "Евгений",
			Patronomic: "Викторович",
		},
		{
			ClinetID:   "50faa486-8e73-4c31-b10f-c7f24c115cda",
			FamilyName: "Шмельцер",
			Name:       "Вячеслав",
			Patronomic: "Николаевич",
		},
		{
			ClinetID:   "60faa486-8e73-4c31-b10f-c7f24c115cda",
			FamilyName: "Виевская",
			Name:       "Анастасия",
			Patronomic: "Федоровна",
		},
	}
	data, err := json.Marshal(&clients)
	if err != nil {
		t.Fatal(err)
	}
	return data, clients
}

// func TestRequest(t *testing.T) ([]byte, []*requestGetNamesByID) {
// 	t.Helper()
// 	treq := []*requestGetNamesByID{
// 		{
// 			ClientID: "48faa486-8e73-4c31-b10f-c7f24c115cda",
// 		},
// 		{
// 			ClientID: "50faa486-8e73-4c31-b10f-c7f24c115cda",
// 		},
// 		{
// 			ClientID: "60faa486-8e73-4c31-b10f-c7f24c115cda",
// 		},
// 	}
// 	data, err := json.Marshal(&treq)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	return data, treq
// }

func TestNewHTTPClient(t *testing.T) *HTTPClient {
	t.Helper()

	return &HTTPClient{
		baseURL: "http://localhost",
		userAgent: "go client",
		client: &http.Client{},
	}
}

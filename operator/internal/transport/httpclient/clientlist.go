package httpclient

import (
	"encoding/json"
	"net/http"

	"github.com/fgituser/management-client-psychologist.services/operator/internal/model"
	"github.com/pkg/errors"
)

type responseClientsList struct {
	ID           string `json:"id,omitempty"`
	FamilyName   string `json:"family_name,omitempty"`
	Name         string `json:"name,omitempty"`
	Patronomic   string `json:"patronomic,omitempty"`
	Psychologist struct {
		ID string `json:"id"`
	} `json:"psychologist"`
}

//ClientsList get clients
func (h *HTTPClient) ClientsList() ([]*model.Client, error) {
	res, err := h.Do(nil, http.MethodGet, "/client/list", userRole)
	if err != nil {
		return nil, errors.Wrap(err, "an error occured get clients")
	}

	clients := make([]*responseClientsList, 0)
	if err := json.Unmarshal(res, &clients); err != nil {
		return nil, errors.Wrap(err, "an error occured get client lesson list")
	}
	return responseClientsListToModelClient(clients), nil
}

func responseClientsListToModelClient(res []*responseClientsList) []*model.Client {
	clients := make([]*model.Client, 0)
	for _, r := range res {
		clients = append(clients, &model.Client{
			ID:         r.ID,
			FamilyName: r.FamilyName,
			Name:       r.Name,
			Patronomic: r.Patronomic,
			Psychologist: &model.Psychologist{
				ID: r.Psychologist.ID,
			},
		})
	}
	return clients
}

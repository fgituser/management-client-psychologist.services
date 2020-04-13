package httpclient

import (
	"encoding/json"
	"net/http"

	"github.com/fgituser/management-client-psychologist.services/operator/internal/model"
	"github.com/pkg/errors"
)

type responsePsychologistListByID struct {
	ID         string `json:"id,omitempty"`
	FamilyName string `json:"family_name,omitempty"`
	Name       string `json:"name,omitempty"`
	Patronomic string `json:"patronomic,omitempty"`
}

//PsychologistListByID get psychologist names from id
func (h *HTTPClient) PsychologistListByID(psychologist []*model.Psychologist) ([]*model.Psychologist, error) {
	body, err := json.Marshal(psychologist)
	if err != nil {
		return nil, errors.Wrap(err, "an error accured while get psychologist by id")
	}
	res, err := h.Do(body, http.MethodPost, "/employees/list_by_id", userRole)
	if err != nil {
		return nil, errors.Wrap(err, "an error accured while get psychologist by id")
	}

	psc := make([]*responsePsychologistListByID, 0)
	if err := json.Unmarshal(res, &psc); err != nil {
		return nil, errors.Wrap(err, "an error accured while get psychologist by id")
	}
	return responsePsychologistListByIDToModelPsychologist(psc), nil
}

func responsePsychologistListByIDToModelPsychologist(res []*responsePsychologistListByID) []*model.Psychologist {
	psychologist := make([]*model.Psychologist, 0)
	for _, r := range res {
		psychologist = append(psychologist, &model.Psychologist{
			ID:         r.ID,
			FamilyName: r.FamilyName,
			Name:       r.Name,
			Patronomic: r.Patronomic,
		})
	}
	return psychologist
}

type responsePsychologistList struct {
	ID         string `json:"id,omitempty"`
	FamilyName string `json:"family_name,omitempty"`
	Name       string `json:"name,omitempty"`
	Patronomic string `json:"patronomic,omitempty"`
	Clients    []struct {
		ID string `json:"id"`
	} `json:"clients"`
}

//PsychologistList get all psychologist
func (h *HTTPClient) PsychologistList() ([]*model.Psychologist, error) {
	res, err := h.Do(nil, http.MethodGet, "/employees/list", userRole)
	if err != nil {
		return nil, errors.Wrap(err, "an error accured while get all psychologist")
	}

	pch := make([]*responsePsychologistList, 0)
	if err := json.Unmarshal(res, &pch); err != nil {
		return nil, errors.Wrap(err, "an error accured while get all psychologist")
	}
	return responsePsychologistListToModelPsychologist(pch), nil
}

func responsePsychologistListToModelPsychologist(res []*responsePsychologistList) []*model.Psychologist {
	psychologist := make([]*model.Psychologist, 0)
	for index, r := range res {
		psychologist = append(psychologist, &model.Psychologist{
			ID:         r.ID,
			FamilyName: r.FamilyName,
			Name:       r.Name,
			Patronomic: r.Patronomic,
		})
		for _, c := range r.Clients {
			psychologist[index].Clients = append(psychologist[index].Clients, &model.Client{ID: c.ID})
		}
	}
	return psychologist
}

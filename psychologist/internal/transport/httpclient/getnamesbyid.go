package httpclient

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/model"
	"github.com/pkg/errors"
)

type requestGetNamesByID struct {
	ClientID string `json:"client_id"`
}

type responseGetNamesByID struct {
	ClinetID   string `json:"client_id"`
	FamilyName string `json:"family_name"`
	Name       string `json:"name"`
	Patronomic string `json:"patronomic"`
}

//GetNamesByID getting client names by identifiers
func (h *HTTPClient) GetNamesByID(c []*model.Client) ([]*model.Client, error) {
	payload, err := encodeGetNamesByID(c)
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while getting customer names by identifier")
	}

	dd, err := decodeGetNameByID(res)
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while getting customer names by identifier")
	}

	return convertresponseGetNamesByIDToModelClient(dd, c), nil
}

//Do ...
func (h *HTTPClient) Do(data []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, h.url, bytes.NewBuffer(data))
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while getting customer names by identifier")
	}

	for k, v := range h.headers {
		req.Header.Add(k, v)
	}

	res, err := h.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while getting customer names by identifier")
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err)
	}
}

func convertresponseGetNamesByIDToModelClient(resclient []*responseGetNamesByID, c []*model.Client) []*model.Client {
	nClients := make([]*model.Client, 1)
	for _, cc := range c {
		for _, d := range resclient {
			if cc.ID == d.ClinetID {
				nClients = append(nClients, &model.Client{
					ID:         d.ClinetID,
					FamilyName: d.FamilyName,
					Name:       d.Name,
					Patronomic: d.Patronomic,
				})
				continue
			}
		}
	}
	return nClients
}

func decodeGetNameByID(res *http.Response) ([]*responseGetNamesByID, error) {
	payload := make([]*responseGetNamesByID, 1)
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		return nil, errors.Wrap(err, "an error occurred while decode get name by id")
	}
	return payload, nil
}

func encodeGetNamesByID(c []*model.Client) ([]byte, error) {
	r := make([]*requestGetNamesByID, 1)
	for _, cc := range c {
		r = append(r, &requestGetNamesByID{ClientID: cc.ID})
	}
	return json.Marshal(r)
}

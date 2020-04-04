package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

//TODO: check user role
const (
	rolePsychologist = "psichologist"
)

//GetNamesByID getting client names by identifiers
func (h *HTTPClient) GetNamesByID(c []*model.Client, psychologistID, userRole string) ([]*model.Client, error) {

	payload, err := encodeGetNamesByIDToRequest(c)
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while getting customer names by identifier")
	}

	res, err := h.Do(payload,
		fmt.Sprintf("/clients/psychologist/%v/names", psychologistID),
		userRole)
	if err != nil {
		return nil, errors.Wrap(err, "an error occured hole get name by id")
	}

	dd, err := decodeGetNameByID(res)
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while getting customer names by identifier")
	}

	return convertresponseGetNamesByIDToModelClient(dd, c), nil
}

//Do send request
func (h *HTTPClient) Do(data []byte, url string, hrole string) ([]byte, error) {

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v%v", h.baseURL, url), bytes.NewBuffer(data))
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while getting customer names by identifier")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", h.userAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Role", hrole)

	res, err := h.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while getting customer names by identifier")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(err, "an error occured while send request, server return response code %v", res.StatusCode)
	}
	return ioutil.ReadAll(res.Body)
}

func convertresponseGetNamesByIDToModelClient(resclient []*responseGetNamesByID, c []*model.Client) []*model.Client {
	nClients := make([]*model.Client, 0)
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

func decodeGetNameByID(data []byte) ([]*responseGetNamesByID, error) {
	payload := make([]*responseGetNamesByID, 1)
	if err := json.NewDecoder(bytes.NewBuffer(data)).Decode(&payload); err != nil {
		return nil, errors.Wrap(err, "an error occurred while decode get name by id")
	}
	return payload, nil
}

func encodeGetNamesByIDToRequest(c []*model.Client) ([]byte, error) {
	if c == nil {
		return nil, errors.New("an nil params accured while encode get names by id to request")
	}
	r := make([]*requestGetNamesByID, 0)
	for _, cc := range c {
		r = append(r, &requestGetNamesByID{ClientID: cc.ID})
	}
	return json.Marshal(r)
}

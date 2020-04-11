package httpclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

//HTTPClient ...
type HTTPClient struct {
	baseURL   string
	userAgent string
	client    *http.Client
}

//New return new http client
func New(baseURL, userAgent string, client *http.Client) (*HTTPClient, error) {
	if strings.TrimSpace(baseURL) == "" {
		return nil, errors.New("not valid url")
	}

	if client == nil {
		return nil, errors.New("an error accurred New httpclient: param client is nil")
	}

	return &HTTPClient{
		baseURL:   baseURL,
		userAgent: userAgent,
		client:    client,
	}, nil
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

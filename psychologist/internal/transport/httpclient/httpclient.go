package httpclient

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

//HTTPClient ...
type HTTPClient struct {
	url     string
	headers map[string]string
	client  *http.Client
}

//New return new http client
func New(url string, headers map[string]string) (*HTTPClient, error) {
	if strings.TrimSpace(url) == "" {
		return nil, errors.New("not valid url")
	}

	return &HTTPClient{
		url:     url,
		headers: headers,
		client:  &http.Client{},
	}, nil
}

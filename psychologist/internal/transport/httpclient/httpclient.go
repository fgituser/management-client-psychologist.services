package httpclient

import (
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

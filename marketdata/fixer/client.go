package fixer

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/http"
	"net/url"
)

type Client struct {
	blockatlas.Request
	APIKey string
}

type Fixer struct {
	client          Client
	BaseURL, APIKey string
}

func (f *Fixer) Init() error {
	f.client = ClientInit(f.BaseURL, f.APIKey)
	return nil
}

func ClientInit(baseUrl, apiKey string) Client {
	return Client{
		Request: blockatlas.Request{
			BaseUrl:    baseUrl,
			HttpClient: http.DefaultClient,
			ErrorHandler: func(res *http.Response, uri string) error {
				return nil
			},
		},
		APIKey: apiKey,
	}
}

func (f *Fixer) fetchLatestRates() (*LatestRatesResponse, error) {
	values := url.Values{
		"access_key": {f.APIKey},
		"base":       {"USD"}, // Base USD supported only in paid api
	}

	latest := new(LatestRatesResponse)
	err := f.client.Get(&latest, "/latest", values)
	if err != nil {
		return nil, err
	}

	return latest, err
}

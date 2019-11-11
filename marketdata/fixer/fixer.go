package fixer

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
	"time"
)

type Fixer struct {
	blockatlas.Request
	APIKey     string
	UpdateTime time.Duration
}

func InitFixer() *Fixer {
	return &Fixer{
		Request:    blockatlas.InitClient(viper.GetString("market.fixer_api")),
		APIKey:     viper.GetString("market.fixer_key"),
		UpdateTime: time.Second * 5,
	}
}

func (f *Fixer) fetchLatestRates() (fixer []blockatlas.Rate, err error) {
	values := url.Values{
		"access_key": {f.APIKey},
		"base":       {"USD"}, // Base USD supported only in paid api
	}
	var latest Latest
	err = f.Get(&latest, "latest", values)
	if err != nil {
		return
	}

	for currency, rate := range latest.Rates {
		fixer = append(fixer, blockatlas.Rate{Currency: currency, Rate: rate, Timestamp: latest.Timestamp})
	}
	return
}

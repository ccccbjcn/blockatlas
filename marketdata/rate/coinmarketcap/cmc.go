package cmc

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/marketdata/rate"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
	"time"
)

type Cmc struct {
	rate.Rate
}

func InitRate() rate.Provider {
	cmc := &Cmc{
		Rate: rate.Rate{
			Id:         "cmc",
			Request:    blockatlas.InitClient(viper.GetString("market.cmc_api")),
			UpdateTime: time.Second * 30,
		},
	}
	cmc.Headers["X-CMC_PRO_API_KEY"] = viper.GetString("market.cmc_api_key")
	return cmc
}

func (c *Cmc) FetchLatestRates() (rates blockatlas.Rates, err error) {
	var prices CoinPrices
	err = c.Get(&prices, "v1/cryptocurrency/listings/latest", url.Values{"limit": {"5000"}, "convert": {"USD"}})
	if err != nil {
		return
	}
	rates = normalizeRates(prices)
	return
}

func normalizeRates(prices CoinPrices) (rates blockatlas.Rates) {
	for _, price := range prices.Data {
		if price.Platform != nil {
			continue
		}
		rates = append(rates, blockatlas.Rate{
			Currency:  price.Symbol,
			Rate:      1.0 / price.Quote.USD.Price,
			Timestamp: price.LastUpdated.Unix(),
		})
	}
	return
}

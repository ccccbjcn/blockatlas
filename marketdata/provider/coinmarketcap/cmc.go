package cmc

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/marketdata/provider"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/url"
	"time"
)

type Market struct {
	provider.Market
}

func InitMarket() provider.Provider {
	m := &Market{
		Market: provider.Market{
			Id:         "cmc",
			Name:       "CoinMarketCap",
			URL:        "https://coinmarketcap.com/",
			Request:    blockatlas.InitClient("https://pro-api.coinmarketcap.com"),
			UpdateTime: time.Second * 5,
		},
	}
	m.Headers["X-CMC_PRO_API_KEY"] = viper.GetString("market.cmc_api_key")
	return m
}

func (p *Market) GetData() (blockatlas.Tickers, error) {
	var prices CoinPrices
	err := p.Get(&prices, "v1/cryptocurrency/listings/latest", url.Values{"limit": {"5000"}, "convert": {"USD"}})
	if err != nil {
		return nil, err
	}
	return normalizeTickers(prices), nil
}

func normalizeTicker(price Data) (*blockatlas.Ticker, error) {
	value24h := percentageChange(price.Quote.USD.Price, price.Quote.USD.PercentChange24h)

	tokenId := ""
	symbol := price.Symbol
	coinType := blockatlas.TypeCoin
	if price.Platform != nil {
		coinType = blockatlas.TypeToken
		symbol = price.Platform.Symbol
		tokenId = price.Symbol
	}

	return &blockatlas.Ticker{
		Coin:     symbol,
		CoinType: coinType,
		TokenId:  tokenId,
		Price: blockatlas.TickerPrice{
			Value:     price.Quote.USD.Price,
			Change24h: value24h,
		},
		LastUpdate: time.Now(),
	}, nil
}

func normalizeTickers(prices CoinPrices) (tickers blockatlas.Tickers) {
	for _, price := range prices.Data {
		t, err := normalizeTicker(price)
		if err != nil {
			logger.Error(err)
			continue
		}
		tickers = append(tickers, *t)
	}
	return
}

func percentageChange(value, percent float64) float64 {
	return value * (percent / 100)
}

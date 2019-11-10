package cmc

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/marketdata/provider"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"time"
)

type Market struct {
	provider.Market
}

func InitMarket() provider.Provider {
	api := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?limit=1000&convert=BTC"
	m := &Market{
		Market: provider.Market{
			Id:         "cmc",
			Name:       "CoinMarketCap",
			URL:        "https://coinmarketcap.com/",
			Request:    blockatlas.InitClient(api),
			UpdateTime: time.Second * 1,
		},
	}
	m.Headers["X-CMC_PRO_API_KEY"] = viper.GetString("market.cmc_api_key")
	return m
}

func (p *Market) GetData() ([]blockatlas.Ticker, error) {
	var prices CoinPrices
	err := p.Get(&prices, "", nil)
	if err != nil {
		return nil, err
	}
	return NormalizeTickers(prices), nil
}

func NormalizeTicker(price Data) (*blockatlas.Ticker, error) {
	value24h := percentageChange(price.Quote.BTC.Price, price.Quote.BTC.PercentChange24h)

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
			Value:     price.Quote.BTC.Price,
			Change24h: value24h,
		},
		LastUpdate: time.Now(),
	}, nil
}

func NormalizeTickers(prices CoinPrices) (tickers []blockatlas.Ticker) {
	for _, price := range prices.Data {
		t, err := NormalizeTicker(price)
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

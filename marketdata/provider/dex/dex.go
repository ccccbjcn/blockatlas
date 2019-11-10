package dex

import (
	"github.com/trustwallet/blockatlas/marketdata/provider"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"strconv"
	"time"
)

type CoinPrice struct {
	Symbol         string `json:"symbol"`
	BaseAssetName  string `json:"baseAssetName"`
	QuoteAssetName string `json:"quoteAssetName"`
	PriceChange    string `json:"priceChange"`
	LastPrice      string `json:"lastPrice"`
}

type Market struct {
	provider.Market
}

func InitMarket() provider.Provider {
	api := "https://dex.binance.org/api/v1/ticker/24hr?limit=1000"
	m := &Market{
		Market: provider.Market{
			Id:         "dex",
			Name:       "Binance Dex",
			URL:        "https://www.binance.org/",
			Request:    blockatlas.InitClient(api),
			UpdateTime: time.Second * 10,
		},
	}
	return m
}

func (p *Market) GetData() ([]blockatlas.Ticker, error) {
	var prices []CoinPrice
	err := p.Get(&prices, "", nil)
	if err != nil {
		return nil, err
	}
	return NormalizeTickers(prices), nil
}

func NormalizeTicker(price CoinPrice) (*blockatlas.Ticker, error) {
	value, err := strconv.ParseFloat(price.LastPrice, 64)
	if err != nil {
		return nil, errors.E(err, "NormalizeTicker parse value error",
			errors.Params{"LastPrice": price.LastPrice, "Symbol": price.Symbol})
	}
	value24h, err := strconv.ParseFloat(price.PriceChange, 64)
	if err != nil {
		return nil, errors.E(err, "NormalizeTicker parse value24h error",
			errors.Params{"PriceChange": price.PriceChange, "Symbol": price.Symbol})
	}
	return &blockatlas.Ticker{
		Coin:     price.Symbol,
		CoinType: blockatlas.TypeCoin,
		Price: blockatlas.TickerPrice{
			Value:     value,
			Change24h: value24h,
		},
		LastUpdate: time.Now(),
	}, nil
}

func NormalizeTickers(prices []CoinPrice) (tickers []blockatlas.Ticker) {
	for _, price := range prices {
		t, err := NormalizeTicker(price)
		if err != nil {
			logger.Error(err)
			continue
		}
		tickers = append(tickers, *t)
	}
	return
}

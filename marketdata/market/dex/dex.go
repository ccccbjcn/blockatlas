package dex

import (
	"github.com/trustwallet/blockatlas/marketdata/market"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/url"
	"strconv"
	"time"
)

type Market struct {
	market.Market
}

func InitMarket() market.Provider {
	m := &Market{
		Market: market.Market{
			Id:         "dex",
			Name:       "Binance Dex",
			URL:        "https://www.binance.org/",
			Request:    blockatlas.InitClient("https://dex.binance.org/api"),
			UpdateTime: time.Second * 6,
		},
	}
	return m
}

func (p *Market) GetData() (blockatlas.Tickers, error) {
	var prices []CoinPrice
	err := p.Get(&prices, "v1/ticker/24hr", url.Values{"limit": {"1000"}})
	if err != nil {
		return nil, err
	}
	return normalizeTickers(prices), nil
}

func normalizeTicker(price CoinPrice) (*blockatlas.Ticker, error) {
	value, err := strconv.ParseFloat(price.LastPrice, 64)
	if err != nil {
		return nil, errors.E(err, "normalizeTicker parse value error",
			errors.Params{"LastPrice": price.LastPrice, "Symbol": price.Symbol})
	}
	value24h, err := strconv.ParseFloat(price.PriceChange, 64)
	if err != nil {
		return nil, errors.E(err, "normalizeTicker parse value24h error",
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

func normalizeTickers(prices []CoinPrice) (tickers blockatlas.Tickers) {
	for _, price := range prices {
		t, err := normalizeTicker(price)
		if err != nil {
			logger.Error(err)
			continue
		}
		tickers = append(tickers, *t)
	}
	return
}

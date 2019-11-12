package dex

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/marketdata/market"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/url"
	"strconv"
	"time"
)

const (
	quoteAsset = "BNB"
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
			Request:    blockatlas.InitClient(viper.GetString("market.dex_api")),
			UpdateTime: time.Second * 1,
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
	rate, err := p.Storage.GetRate(quoteAsset)
	if err != nil {
		return nil, errors.E(err, "rate not found", logger.Params{"asset": quoteAsset})
	}
	result := normalizeTickers(prices)
	result.ApplyRate(1.0/rate.Rate, "USD")
	return result, nil
}

func normalizeTicker(price CoinPrice) (*blockatlas.Ticker, error) {
	if price.QuoteAssetName != quoteAsset {
		return nil, errors.E("invalid quote asset",
			errors.Params{"Symbol": price.BaseAssetName, "QuoteAsset": price.QuoteAssetName})
	}
	value, err := strconv.ParseFloat(price.LastPrice, 64)
	if err != nil {
		return nil, errors.E(err, "normalizeTicker parse value error",
			errors.Params{"LastPrice": price.LastPrice, "Symbol": price.BaseAssetName})
	}
	value24h, err := strconv.ParseFloat(price.PriceChange, 64)
	if err != nil {
		return nil, errors.E(err, "normalizeTicker parse value24h error",
			errors.Params{"PriceChange": price.PriceChange, "Symbol": price.BaseAssetName})
	}
	return &blockatlas.Ticker{
		Coin:     price.BaseAssetName,
		CoinType: blockatlas.TypeCoin,
		Price: blockatlas.TickerPrice{
			Value:     value,
			Change24h: value24h,
			Currency:  "BNB",
		},
		LastUpdate: time.Now(),
	}, nil
}

func normalizeTickers(prices []CoinPrice) (tickers blockatlas.Tickers) {
	for _, price := range prices {
		t, err := normalizeTicker(price)
		if err != nil {
			continue
		}
		tickers = append(tickers, t)
	}
	return
}

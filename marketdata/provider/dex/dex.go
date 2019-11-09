package dex

import (
	"github.com/trustwallet/blockatlas/marketdata/provider"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/storage"
	"time"
)

type CoinPrice struct {
	Symbol             string `json:"symbol"`
	BaseAssetName      string `json:"baseAssetName"`
	QuoteAssetName     string `json:"quoteAssetName"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQuantity       string `json:"lastQuantity"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	OpenTime           uint64 `json:"openTime"`
	CloseTime          uint64 `json:"closeTime"`
	FirstId            string `json:"firstId"`
	LastId             string `json:"lastId"`
	BidPrice           string `json:"bidPrice"`
	BidQuantity        string `json:"bidQuantity"`
	AskPrice           string `json:"askPrice"`
	AskQuantity        string `json:"askQuantity"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	Count              uint64 `json:"count"`
}

type Market struct {
	provider.Market
}

func InitMarket(storage storage.Market) *Market {
	m := &Market{
		Market: provider.Market{
			Id:      "dex",
			Name:    "Binance Dex",
			URL:     "https://www.binance.org/",
			Api:     "https://dex.binance.org/api/v1/ticker/24hr?limit=1000",
			Storage: storage,
		},
	}
	m.Market.GetData = m.GetData
	return m
}

func (p *Market) GetData() ([]blockatlas.Ticker, error) {
	var prices []CoinPrice
	p.Client.Get(&prices, "", nil)
	ticker := []blockatlas.Ticker{
		{Coin: "BTC", CoinType: blockatlas.TypeCoin, Price: blockatlas.TickerPrice{Value: 111, Change24h: float64(time.Now().Unix())}, LastUpdate: time.Now()},
		{Coin: "ETH", TokenId: "HT", CoinType: blockatlas.TypeToken, Price: blockatlas.TickerPrice{Value: 222, Change24h: float64(time.Now().Unix())}, LastUpdate: time.Now()},
		{Coin: "OMNI", TokenId: "USDT", CoinType: blockatlas.TypeToken, Price: blockatlas.TickerPrice{Value: 333, Change24h: float64(time.Now().Unix())}, LastUpdate: time.Now()},
		{Coin: "ETH", TokenId: "BLA", CoinType: blockatlas.TypeToken, Price: blockatlas.TickerPrice{Value: 444, Change24h: float64(time.Now().Unix())}, LastUpdate: time.Now()},
	}
	return ticker, nil
}

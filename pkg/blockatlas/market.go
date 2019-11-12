package blockatlas

import (
	"time"
)

const (
	TypeCoin  CoinType = "coin"
	TypeToken CoinType = "token"
)

type CoinType string

type TickerResponse struct {
	Currency string  `json:"currency"`
	Result   Tickers `json:"result"`
}

type Ticker struct {
	Coin       string      `json:"coin"`
	TokenId    string      `json:"token_id,omitempty"`
	CoinType   CoinType    `json:"type"`
	Price      TickerPrice `json:"price"`
	LastUpdate time.Time   `json:"last_update"`
	Error      string      `json:"error,omitempty"`
}

type TickerPrice struct {
	Value     float64 `json:"value"`
	Change24h float64 `json:"change_24h"`
	Currency  string  `json:"currency"`
	Provider  string  `json:"provider"`
}

type Rate struct {
	Currency  string  `json:"currency"`
	Rate      float64 `json:"rate"`
	Timestamp int64   `json:"timestamp"`
}

type Rates []Rate
type Tickers []*Ticker

func (ts Tickers) ApplyRate(rate float64, currency string) {
	for _, t := range ts {
		t.ApplyRate(rate, currency)
	}
}

func (t *Ticker) ApplyRate(rate float64, currency string) {
	if t.Price.Currency == currency {
		return
	}
	t.Price.Value *= rate
	t.Price.Currency = currency
}

package blockatlas

import (
	"time"
)

const (
	TypeCoin  CoinType = "coin"
	TypeToken CoinType = "token"
)

type CoinType string

type Ticker struct {
	Coin       string      `json:"coin"`
	TokenId    string      `json:"token_id,omitempty"`
	CoinType   CoinType    `json:"type"`
	Price      TickerPrice `json:"price"`
	LastUpdate time.Time   `json:"last_update"`
}

type TickerPrice struct {
	Value     float64 `json:"value"`
	Change24h float64 `json:"change_24h"`
}

type Rate struct {
	Currency  string  `json:"currency"`
	Rate      float64 `json:"rate"`
	Timestamp uint64  `json:"timestamp"`
}

package market

import (
	"time"
)

type Priority int
type Providers map[Priority]Provider

type CoinType string

const (
	TypeCoin  CoinType = "coin"
	TypeToken CoinType = "token"

	defaultUpdateTime = time.Second * 2
)

type Provider struct {
	ID             string
	Name           string
	URL            string
	UpdateTime     time.Duration
	GetData        func() (interface{}, error)
	NormalizeCoins func(interface{}) ([]Ticker, error)
	Storage        Storage
}

type Ticker struct {
	Coin       string    `json:"coin"`
	TokenId    string    `json:"token_id,omitempty"`
	CoinType   CoinType  `json:"type"`
	Price      Price     `json:"price"`
	LastUpdate time.Time `json:"last_update"`
}

type Price struct {
	Value     float64 `json:"value"`
	Change24h float64 `json:"change_24h"`
}

//type Storage interface {
//	SaveTicker(entity string, coin Ticker) error
//	GetTicker(entity, coin, token string) (*Ticker, error)
//	SaveMarketPriority(id string, p Priority) error
//}

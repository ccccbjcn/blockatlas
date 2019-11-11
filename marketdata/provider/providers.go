package provider

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"time"
)

type MarketPriority int
type Providers map[MarketPriority]Provider

type Provider interface {
	Init() error
	GetName() string
	GetId() string
	GetUpdateTime() time.Duration
	GetData() (blockatlas.Tickers, error)
}

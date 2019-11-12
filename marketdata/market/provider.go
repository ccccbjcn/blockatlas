package market

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"time"
)

type Priority int
type Providers map[Priority]Provider

type Provider interface {
	Init() error
	GetName() string
	GetId() string
	GetUpdateTime() time.Duration
	GetData() (blockatlas.Tickers, error)
	GetType() string
}

package market

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/storage"
	"time"
)

type Priority int
type Providers map[Priority]Provider

type Provider interface {
	Init(storage.Market) error
	GetName() string
	GetId() string
	GetUpdateTime() time.Duration
	GetData() (blockatlas.Tickers, error)
	GetType() string
}

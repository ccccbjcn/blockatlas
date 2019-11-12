package rate

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"time"
)

type Provider interface {
	Init() error
	FetchLatestRates() (blockatlas.Rates, error)
	GetUpdateTime() time.Duration
	GetId() string
	GetType() string
}

package provider

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
	"time"
)

const (
	defaultUpdateTime = time.Second * 20
)

type MarketPriority int
type Providers map[MarketPriority]Provider

type Provider interface {
	Init() error
	Run() error
	GetId() string
	GetUpdateTime() time.Duration
}

type Market struct {
	Id         string
	Name       string
	URL        string
	UpdateTime time.Duration
	Storage    storage.Market
	Api        string
	Client     blockatlas.Request
	GetData    func() ([]blockatlas.Ticker, error)
}

func (m *Market) GetId() string {
	return m.Id
}

func (m *Market) GetUpdateTime() time.Duration {
	return m.UpdateTime
}

func (m *Market) Init() error {
	logger.Info("Init Provider", logger.Params{"provider": m.GetId()})
	if m.Storage == nil {
		return errors.E("Provider: Storage cannot be nil")
	}
	if len(m.Api) == 0 {
		return errors.E("Provider: Api cannot be empty")
	}
	if len(m.Id) == 0 {
		return errors.E("Provider: Id cannot be empty")
	}
	if len(m.Name) == 0 {
		return errors.E("Provider: Name cannot be empty")
	}
	if m.UpdateTime == 0 {
		m.UpdateTime = defaultUpdateTime
	}
	m.Client = blockatlas.InitClient(m.Api)
	return nil
}

func (m *Market) Run() error {
	logger.Info("Starting market data task...", logger.Params{"Provider": m.Name, "ProviderId": m.Id})
	data, err := m.GetData()
	if err != nil {
		return errors.E(err, "GetData")
	}
	for _, result := range data {
		err = m.Storage.SaveTicker(m.Id, result)
		if err != nil {
			logger.Error(errors.E(err, "SaveTicker",
				errors.Params{"result": result}))
		}
	}
	logger.Info("Market data result", logger.Params{"markets": len(data)})
	return nil
}

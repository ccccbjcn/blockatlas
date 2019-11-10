package provider

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"time"
)

const (
	defaultUpdateTime = time.Second * 20
)

type Market struct {
	blockatlas.Request
	Id         string
	Name       string
	URL        string
	UpdateTime time.Duration
}

func (m *Market) GetName() string {
	return m.Name
}

func (m *Market) GetId() string {
	return m.Id
}

func (m *Market) GetUpdateTime() time.Duration {
	return m.UpdateTime
}

func (m *Market) Init() error {
	logger.Info("Init Provider", logger.Params{"provider": m.GetId()})
	if len(m.Id) == 0 {
		return errors.E("Provider: Id cannot be empty")
	}
	if len(m.Name) == 0 {
		return errors.E("Provider: Name cannot be empty")
	}
	if m.UpdateTime == 0 {
		m.UpdateTime = defaultUpdateTime
	}
	return nil
}

package market

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
)

func (md *Provider) Init() error {
	if md.Storage == nil {
		return errors.E("Provider: Storage cannot be nil")
	}
	if len(md.ID) == 0 {
		return errors.E("Provider: ID cannot be empty")
	}
	if len(md.Name) == 0 {
		return errors.E("Provider: Name cannot be empty")
	}
	if md.GetData == nil {
		return errors.E("Provider: GetData cannot be nil")
	}
	if md.NormalizeCoins == nil {
		return errors.E("Provider: NormalizeCoins cannot be nil")
	}
	if md.UpdateTime == 0 {
		md.UpdateTime = defaultUpdateTime
	}
	return nil
}

func (md *Provider) Run() error {
	data, err := md.GetData()
	if err != nil {
		return errors.E(err, "GetData")
	}
	markets, err := md.NormalizeCoins(data)
	if err != nil {
		return errors.E(err, "NormalizeCoins", errors.Params{"data": data})
	}
	for _, result := range markets {
		err = md.Storage.SaveTicker(md.ID, result)
		if err != nil {
			logger.Error(errors.E(err, "SaveTicker",
				errors.Params{"result": result}))
		}
	}
	return nil
}

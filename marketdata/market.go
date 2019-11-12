package marketdata

import (
	"github.com/robfig/cron"
	"github.com/trustwallet/blockatlas/marketdata/market"
	cmc "github.com/trustwallet/blockatlas/marketdata/market/coinmarketcap"
	"github.com/trustwallet/blockatlas/marketdata/market/dex"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
)

func InitMarkets(storage storage.Market) {
	addMarkets(storage,
		market.Providers{
			0: dex.InitMarket(),
			1: cmc.InitMarket(),
		})
}

func addMarkets(storage storage.Market, ps market.Providers) {
	c := cron.New()
	priorityList := make(map[int]string)
	for priority, p := range ps {
		scheduleTasks(storage, p, c)
		priorityList[int(priority)] = p.GetId()
	}
	err := storage.SaveMarketPriority(priorityList)
	if err != nil {
		logger.Error(err, "SaveMarketPriority", logger.Params{"priorityList": priorityList})
	}
	c.Start()
}

func runMarket(storage storage.Market, p market.Provider) error {
	data, err := p.GetData()
	if err != nil {
		return errors.E(err, "GetData")
	}
	for _, result := range data {
		err = storage.SaveTicker(p.GetId(), result)
		if err != nil {
			logger.Error(errors.E(err, "SaveTicker",
				errors.Params{"result": result}))
		}
	}
	logger.Info("Market data result", logger.Params{"markets": len(data)})
	return nil
}

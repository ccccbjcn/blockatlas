package market

import (
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/robfig/cron"
	"github.com/trustwallet/blockatlas/marketdata/provider"
	cmc "github.com/trustwallet/blockatlas/marketdata/provider/coinmarketcap"
	"github.com/trustwallet/blockatlas/marketdata/provider/dex"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
	"time"
)

const (
	backoffValue = 2
)

func InitProviders(storage storage.Market) {
	addManyMarketData(storage,
		provider.Providers{
			0: dex.InitMarket(),
			1: cmc.InitMarket(),
		})
}

func addManyMarketData(storage storage.Market, ps provider.Providers) {
	c := cron.New()
	priorityList := make(map[int]string)
	for priority, p := range ps {
		scheduleRun(storage, p, c)
		priorityList[int(priority)] = p.GetId()
	}
	err := storage.SaveMarketPriority(priorityList)
	if err != nil {
		logger.Error(err, "SaveMarketPriority", logger.Params{"priorityList": priorityList})
	}
	c.Start()
}

func scheduleRun(storage storage.Market, p provider.Provider, c *cron.Cron) {
	err := p.Init()
	if err != nil {
		logger.Error(err, "Init Provider Error", logger.Params{"provider": p.GetId()})
		return
	}
	t := p.GetUpdateTime().Seconds()
	spec := fmt.Sprintf("@every %ds", uint64(t))
	err = c.AddFunc(spec, func() {
		processBackoff(storage, p)
	})
	if err != nil {
		logger.Error(err, "AddFunc")
	}
}

// processBackoff make a exponential backoff for market run
// errors, increasing the retry in a exponential period for each attempt.
func processBackoff(storage storage.Market, p provider.Provider) {
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = backoffValue * time.Minute
	r := func() error {
		return run(storage, p)
	}

	n := func(err error, t time.Duration) {
		logger.Error(err, "process backoff market", logger.Params{"Duration": t.String()})
	}
	err := backoff.RetryNotify(r, b, n)
	if err != nil {
		logger.Error(err, "Market ProcessBackoff")
	}
}

func run(storage storage.Market, p provider.Provider) error {
	logger.Info("Starting market data task...", logger.Params{"Provider": p.GetName(), "ProviderId": p.GetId()})
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

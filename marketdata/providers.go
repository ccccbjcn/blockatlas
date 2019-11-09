package market

import (
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/robfig/cron"
	"github.com/trustwallet/blockatlas/marketdata/provider"
	"github.com/trustwallet/blockatlas/marketdata/provider/dex"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
	"time"
)

const (
	backoffValue = 3
)

func InitProviders(storage storage.Market) {
	AddManyMarketData(storage,
		provider.Providers{
			0: dex.InitMarket(storage),
			//1: provider.Provider{
			//	Id:   "cmc",
			//	Name: "CoinMarketCap",
			//	URL:  "https://coinmarketcap.com/",
			//	UpdateTime: time.Second,
			//},
		})

}

func AddManyMarketData(storage storage.Market, ps provider.Providers) {
	c := cron.New()
	priorityList := make(map[int]string)
	for priority, p := range ps {
		ScheduleRun(p, c)
		priorityList[int(priority)] = p.GetId()
	}
	err := storage.SaveMarketPriority(priorityList)
	if err != nil {
		logger.Error(err, "SaveMarketPriority", logger.Params{"priorityList": priorityList})
	}
	c.Start()
	<-make(chan bool)
}

func ScheduleRun(m provider.Provider, c *cron.Cron) {
	err := m.Init()
	if err != nil {
		logger.Error(err, "Init Provider Error", logger.Params{"provider": m.GetId()})
		return
	}
	t := m.GetUpdateTime().Seconds()
	spec := fmt.Sprintf("@every %ds", uint64(t))
	err = c.AddFunc(spec, func() {
		ProcessBackoff(m.Run)
	})
	if err != nil {
		logger.Error(err, "AddFunc")
	}
}

// processBackoff make a exponential backoff for market run
// errors, increasing the retry in a exponential period for each attempt.
func ProcessBackoff(handler func() error) {
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = backoffValue * time.Minute
	r := func() error {
		return handler()
	}

	n := func(err error, t time.Duration) {
		logger.Error(err, "process Backoff", logger.Params{"Duration": t.String()})
	}
	err := backoff.RetryNotify(r, b, n)
	if err != nil {
		logger.Error(err, "ProcessBackoff")
	}
}

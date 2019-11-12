package marketdata

import (
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/robfig/cron"
	"github.com/trustwallet/blockatlas/marketdata/market"
	"github.com/trustwallet/blockatlas/marketdata/rate"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
	"time"
)

const (
	backoffValue = 2
)

type Provider interface {
	Init(storage.Market) error
	GetId() string
	GetType() string
	GetUpdateTime() time.Duration
}

// processBackoff make a exponential backoff for market run
// errors, increasing the retry in a exponential period for each attempt.
func processBackoff(storage storage.Market, md Provider) {
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = backoffValue * time.Minute
	r := func() error {
		return run(storage, md)
	}

	n := func(err error, t time.Duration) {
		logger.Error(err, "process backoff market", logger.Params{"Duration": t.String()})
	}
	err := backoff.RetryNotify(r, b, n)
	if err != nil {
		logger.Error(err, "Market ProcessBackoff")
	}
}

func scheduleTasks(storage storage.Market, md Provider, c *cron.Cron) {
	err := md.Init(storage)
	if err != nil {
		logger.Error(err, "Init Market Error", logger.Params{"Type": md.GetType(), "Market": md.GetId()})
		return
	}
	t := md.GetUpdateTime().Seconds()
	spec := fmt.Sprintf("@every %ds", uint64(t))
	err = c.AddFunc(spec, func() {
		processBackoff(storage, md)
	})
	processBackoff(storage, md)
	if err != nil {
		logger.Error(err, "AddFunc")
	}
}

func run(storage storage.Market, md Provider) error {
	logger.Info("Starting market data task...", logger.Params{"Type": md.GetType(), "Market": md.GetId()})
	switch m := md.(type) {
	case market.Provider:
		return runMarket(storage, m)
	case rate.Provider:
		return runRate(storage, m)
	}
	return errors.E("invalid market interface")
}

package fixer

import (
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/robfig/cron"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
	"time"
)

const (
	backoffValue = 2
)

func Start(storage storage.Rates) error {
	f := initFixer()
	c := cron.New()
	t := f.UpdateTime.Seconds()
	spec := fmt.Sprintf("@every %ds", uint64(t))
	err := c.AddFunc(spec, func() {
		processBackoff(storage, f)
	})
	if err != nil {
		return err
	}
	c.Start()
	return nil
}

func processBackoff(storage storage.Rates, f *Fixer) {
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = backoffValue * time.Minute
	r := func() error {
		logger.Info("Updating fiat rates ...")
		rates, err := f.fetchLatestRates()
		if err != nil {
			return err
		}
		storage.SaveRates(rates)
		return nil
	}

	n := func(err error, t time.Duration) {
		logger.Error(err, "process backoff fixer", logger.Params{"Duration": t.String()})
	}
	err := backoff.RetryNotify(r, b, n)
	if err != nil {
		logger.Error(err, "Fixer ProcessBackoff")
	}
}

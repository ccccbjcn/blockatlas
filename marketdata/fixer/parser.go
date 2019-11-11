package fixer

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/logger"
)

type Indexer interface {
	Init() error
}

func Start() {
	f := Fixer{
		BaseURL: viper.GetString("fixer.api"),
		APIKey:  viper.GetString("fixer.api_key"),
	}
	f.Init()

	logger.Info("Updating fiat rates ...")
	rates, err := f.fetchLatestRates()
	if err != nil {
		logger.Error(err)
		return
	}

	saveRates(rates)
}

package cmd

import (
	"github.com/spf13/cobra"
	market "github.com/trustwallet/blockatlas/marketdata"
	"github.com/trustwallet/blockatlas/marketdata/fixer"
	"github.com/trustwallet/blockatlas/pkg/logger"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync all markets",
	Run:   syncMarketData,
}

func syncMarketData(cmd *cobra.Command, args []string) {
	market.InitProviders(Storage)
	err := fixer.Start(Storage)
	if err != nil {
		logger.Error(err, "fixer start error")
	}
	<-make(chan bool)
}

func init() {
	rootCmd.AddCommand(syncCmd)
}

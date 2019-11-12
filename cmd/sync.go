package cmd

import (
	"github.com/spf13/cobra"
	"github.com/trustwallet/blockatlas/marketdata"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync all markets",
	Run:   syncMarketData,
}

func syncMarketData(cmd *cobra.Command, args []string) {
	marketdata.InitMarkets(Storage)
	marketdata.InitRates(Storage)
	<-make(chan bool)
}

func init() {
	rootCmd.AddCommand(syncCmd)
}

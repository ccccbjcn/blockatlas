package cmd

import (
	"github.com/spf13/cobra"
	market "github.com/trustwallet/blockatlas/marketdata"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync all markets",
	Run:   syncMarketData,
}

func syncMarketData(cmd *cobra.Command, args []string) {
	market.InitProviders(Storage)

}

func init() {
	rootCmd.AddCommand(syncCmd)
}

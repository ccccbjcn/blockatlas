package dex

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"reflect"
	"testing"
	"time"
)

func TestNormalizeTickers(t *testing.T) {
	tests := []struct {
		name        string
		prices      []CoinPrice
		wantTickers blockatlas.Tickers
	}{
		{
			name:   "",
			prices: []CoinPrice{},
			wantTickers: blockatlas.Tickers{
				{Coin: "BTC", CoinType: blockatlas.TypeCoin, Price: blockatlas.TickerPrice{Value: 111, Change24h: float64(time.Now().Unix())}, LastUpdate: time.Now()},
				{Coin: "ETH", TokenId: "HT", CoinType: blockatlas.TypeToken, Price: blockatlas.TickerPrice{Value: 222, Change24h: float64(time.Now().Unix())}, LastUpdate: time.Now()},
				{Coin: "OMNI", TokenId: "USDT", CoinType: blockatlas.TypeToken, Price: blockatlas.TickerPrice{Value: 333, Change24h: float64(time.Now().Unix())}, LastUpdate: time.Now()},
				{Coin: "ETH", TokenId: "BLA", CoinType: blockatlas.TypeToken, Price: blockatlas.TickerPrice{Value: 444, Change24h: float64(time.Now().Unix())}, LastUpdate: time.Now()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTickers := normalizeTickers(tt.prices); !reflect.DeepEqual(gotTickers, tt.wantTickers) {
				t.Errorf("normalizeTickers() = %v, want %v", gotTickers, tt.wantTickers)
			}
		})
	}
}

package market

import (
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/robfig/cron"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
	"time"
)

const (
	backoffValue      = 3
	defaultUpdateTime = time.Second * 2
)

type Providers map[blockatlas.MarketPriority]Provider

type Provider struct {
	ID             string
	Name           string
	URL            string
	UpdateTime     time.Duration
	GetData        func() (interface{}, error)
	NormalizeCoins func(interface{}) ([]blockatlas.Ticker, error)
	Storage        storage.Market
}

func InitProviders(storage storage.Market) {
	AddManyMarketData(storage,
		Providers{
			0: {
				ID:   "dex",
				Name: "Binance Dex",
				URL:  "https://www.binance.org/",
				GetData: func() (interface{}, error) {
					return "BTC", nil
				},
				NormalizeCoins: func(d interface{}) ([]blockatlas.Ticker, error) {
					return []blockatlas.Ticker{
						{Coin: "BTC", CoinType: blockatlas.TypeCoin, Price: blockatlas.TickerPrice{Value: 555, Change24h: float64(time.Now().Unix())}, LastUpdate: time.Now()},
						{Coin: "ETH", TokenId: "HT", CoinType: blockatlas.TypeToken, Price: blockatlas.TickerPrice{Value: 666, Change24h: float64(time.Now().Unix())}, LastUpdate: time.Now()},
						{Coin: "OMNI", TokenId: "USDT", CoinType: blockatlas.TypeToken, Price: blockatlas.TickerPrice{Value: 777, Change24h: float64(time.Now().Unix())}, LastUpdate: time.Now()},
						{Coin: "OMNI", TokenId: "THT", CoinType: blockatlas.TypeToken, Price: blockatlas.TickerPrice{Value: 888, Change24h: float64(time.Now().Unix())}, LastUpdate: time.Now()},
					}, nil
				},
				Storage:    storage,
				UpdateTime: defaultUpdateTime,
			},
			1: {
				ID:   "cmc",
				Name: "CoinMarketCap",
				URL:  "https://coinmarketcap.com/",
				GetData: func() (interface{}, error) {
					return "BTC", nil
				},
				NormalizeCoins: func(d interface{}) ([]blockatlas.Ticker, error) {
					return []blockatlas.Ticker{
						{Coin: "BTC", CoinType: blockatlas.TypeCoin, Price: blockatlas.TickerPrice{Value: 111, Change24h: float64(time.Now().Unix())}, LastUpdate: time.Now()},
						{Coin: "ETH", TokenId: "HT", CoinType: blockatlas.TypeToken, Price: blockatlas.TickerPrice{Value: 222, Change24h: float64(time.Now().Unix())}, LastUpdate: time.Now()},
						{Coin: "OMNI", TokenId: "USDT", CoinType: blockatlas.TypeToken, Price: blockatlas.TickerPrice{Value: 333, Change24h: float64(time.Now().Unix())}, LastUpdate: time.Now()},
						{Coin: "ETH", TokenId: "BLA", CoinType: blockatlas.TypeToken, Price: blockatlas.TickerPrice{Value: 444, Change24h: float64(time.Now().Unix())}, LastUpdate: time.Now()},
					}, nil
				},
				Storage:    storage,
				UpdateTime: time.Second,
			},
		})

}

func AddManyMarketData(storage storage.Market, ps Providers) {
	c := cron.New()
	priorityList := make(map[blockatlas.MarketPriority]string)
	for priority, provider := range ps {
		ScheduleRun(provider, c)
		priorityList[priority] = provider.ID
	}
	err := storage.SaveMarketPriority(priorityList)
	if err != nil {
		logger.Error(err, "SaveMarketPriority", logger.Params{"priorityList": priorityList})
	}
	c.Start()
	<-make(chan bool)
}

func ScheduleRun(md Provider, c *cron.Cron) {
	t := md.UpdateTime.Seconds()
	spec := fmt.Sprintf("@every %ds", uint64(t))
	err := c.AddFunc(spec, func() {
		ProcessBackoff(md.Run)
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

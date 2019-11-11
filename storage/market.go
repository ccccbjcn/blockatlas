package storage

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"strings"
)

const (
	EntityPriority = "market_priority"
	EntityRates    = "currencies_rates"
)

func (s *Storage) SaveMarketPriority(p map[int]string) error {
	return s.Add(EntityPriority, p)
}

func (s *Storage) GetMarketPriority() (p *map[int]string, err error) {
	err = s.GetValue(EntityPriority, p)
	return
}

func (s *Storage) SaveTicker(entity string, coin blockatlas.Ticker) error {
	cd, err := s.GetTicker(entity, coin.Coin, coin.TokenId)
	if err == nil {
		if cd.LastUpdate.After(coin.LastUpdate) {
			return errors.E("ticker is outdated")
		}
	}
	hm := createHashMap(coin.Coin, coin.TokenId)
	return s.AddHM(entity, hm, coin)
}

func (s *Storage) GetTicker(entity, coin, token string) (blockatlas.Ticker, error) {
	hm := createHashMap(coin, token)
	var cd blockatlas.Ticker
	err := s.GetHMValue(entity, hm, &cd)
	if err != nil {
		return blockatlas.Ticker{}, err
	}
	return cd, nil
}

func (s *Storage) SaveRates(rates []blockatlas.Rate) {
	for _, rate := range rates {
		r, err := s.GetRate(rate.Currency)
		if err == nil && rate.Timestamp < r.Timestamp {
			return
		}
		err = s.AddHM(EntityRates, rate.Currency, &rate)
		if err != nil {
			logger.Error(err, "SaveRates", logger.Params{"rate": rate})
		}
	}
}

func (s *Storage) GetRate(currency string) (rate *blockatlas.Rate, err error) {
	err = s.GetHMValue(EntityRates, currency, &rate)
	return
}

func createHashMap(coin, token string) string {
	if len(token) == 0 {
		return coin
	}
	return strings.Join([]string{coin, token}, "-")
}

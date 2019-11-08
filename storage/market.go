package storage

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strings"
)

const (
	EntityPriority = "market_priority"
)

func (s *Storage) SaveMarketPriority(p map[blockatlas.MarketPriority]string) error {
	return s.Add(EntityPriority, p)
}

func (s *Storage) GetMarketPriority() (p *map[blockatlas.MarketPriority]string, err error) {
	err = s.GetValue(EntityPriority, p)
	return
}

func (s *Storage) SaveTicker(entity string, coin blockatlas.Ticker) error {
	cd, err := s.GetTicker(entity, coin.Coin, coin.TokenId)
	if err == nil && cd != nil {
		if cd.LastUpdate.After(coin.LastUpdate) {
			return err
		}
	}
	hm := createHashMap(coin.Coin, coin.TokenId)
	return s.AddHM(entity, hm, coin)
}

func (s *Storage) GetTicker(entity, coin, token string) (*blockatlas.Ticker, error) {
	hm := createHashMap(coin, token)
	var cd blockatlas.Ticker
	err := s.GetHMValue(entity, hm, &cd)
	if err != nil {
		return nil, err
	}
	return &cd, nil
}

func createHashMap(coin, token string) string {
	if len(token) == 0 {
		return coin
	}
	return strings.Join([]string{coin, token}, "-")
}

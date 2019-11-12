package cmc

import "time"

type CoinPrices struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
	} `json:"status"`
	Data []Data `json:"data"`
}

type Coin struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Slug   string `json:"slug"`
}

type Data struct {
	Coin
	LastUpdated time.Time `json:"last_updated"`
	Platform    *struct {
		Coin
		TokenAddress string `json:"token_address"`
	} `json:"platform"`
	Quote Quote `json:"quote"`
}

type Quote struct {
	USD struct {
		Price            float64 `json:"price"`
		PercentChange24h float64 `json:"percent_change_24h"`
	} `json:"USD"`
}

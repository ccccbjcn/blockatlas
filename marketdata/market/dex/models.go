package dex

type CoinPrice struct {
	Symbol         string `json:"symbol"`
	BaseAssetName  string `json:"baseAssetName"`
	QuoteAssetName string `json:"quoteAssetName"`
	PriceChange    string `json:"priceChange"`
	LastPrice      string `json:"lastPrice"`
}

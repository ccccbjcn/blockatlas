package dex

type CoinPrice struct {
	BaseAssetName  string `json:"baseAssetName"`
	QuoteAssetName string `json:"quoteAssetName"`
	PriceChange    string `json:"priceChange"`
	LastPrice      string `json:"lastPrice"`
}

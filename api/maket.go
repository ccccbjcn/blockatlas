package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/storage"
	"net/http"
	"strconv"
)

type TickerRequest struct {
	Currency string `json:"currency"`
	Market   string `json:"market,omitempty"`
	Assets   []Coin `json:"assets"`
}

type Coin struct {
	Coin     uint                `json:"coin"`
	CoinType blockatlas.CoinType `json:"type"`
	TokenId  string              `json:"token_id,omitempty"`
}

func SetupMarketAPI(router gin.IRouter, db storage.Market) {
	router.Use(ginutils.TokenAuthMiddleware(viper.GetString("market.auth")))
	router.GET("/ticker", getTickerHandler(db))
	router.POST("/ticker", getTickersHandler(db))
}

// @Summary Get ticker value for a specific market
// @Id get_ticker
// @Description Get the ticker value from an market and coin/token
// @Accept json
// @Produce json
// @Tags market
// @Param coin path string true "the coin symbol" default(btc)
// @Param address path string true "the query address" default(tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q)
// @Success 200 {object} api.MarketDataResponse
// @Router /market/v1/ticker [get]
func getTickerHandler(storage storage.Market) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		coinQuery := c.Query("coin")
		coinId, err := strconv.Atoi(coinQuery)
		if err != nil {
			ginutils.RenderError(c, http.StatusInternalServerError, "Invalid coin")
			return
		}
		market := c.DefaultQuery("market", "cmc")
		token := c.Query("token")

		currency := c.DefaultQuery("currency", "USD")
		rate, err := storage.GetRate(currency)
		if err != nil {
			ginutils.RenderError(c, http.StatusInternalServerError, "Invalid currency")
			return
		}

		coinObj := coin.Coins[uint(coinId)]
		result, err := storage.GetTicker(market, coinObj.Symbol, token)
		if err != nil {
			ginutils.RenderError(c, http.StatusInternalServerError, err.Error())
			return
		}
		result.ApplyRate(rate.Rate, currency)
		ginutils.RenderSuccess(c, result)
	}
}

// @Summary Get ticker values for a specific markets
// @Id get_tickers
// @Description Get the ticker values from many markets and coin/token
// @Accept json
// @Produce json
// @Tags ticker
// @Param subscriptions body api.MarketDataRequest true "Ticker"
// @Success 200 {object} api.MarketDataResponse
// @Router /market/v1/tickers [post]
func getTickersHandler(storage storage.Market) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		md := TickerRequest{Currency: "USD", Market: "cmc"}
		if err := c.BindJSON(&md); err != nil {
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}
		rate, err := storage.GetRate(md.Currency)
		if err != nil {
			ginutils.RenderError(c, http.StatusInternalServerError, "Invalid currency")
			return
		}

		tickers := make(blockatlas.Tickers, 0)
		for _, coinRequest := range md.Assets {
			coinObj := coin.Coins[coinRequest.Coin]
			r, err := storage.GetTicker(md.Market, coinObj.Symbol, coinRequest.TokenId)
			if err != nil {
				r.Error = err.Error()
			}
			r.ApplyRate(rate.Rate, md.Currency)
			tickers = append(tickers, r)
		}
		ginutils.RenderSuccess(c, blockatlas.TickerResponse{Currency: md.Currency, Result: tickers})
	}
}

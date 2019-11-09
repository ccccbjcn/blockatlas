package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/storage"
	"net/http"
	"strings"
)

type TickerRequest struct {
	Currency string `json:"currency"`
	Coins    []Coin `json:"coins"`
}

type Coin struct {
	Coin     string              `json:"coin"`
	CoinType blockatlas.CoinType `json:"type"`
	TokenId  string              `json:"token_id,omitempty"`
}

func SetupMarketAPI(router gin.IRouter, db storage.Market) {
	router.Use(ginutils.TokenAuthMiddleware(viper.GetString("market.auth")))
	router.GET("/:coin/ticker", getTickerHandler(db))
	//router.POST("/webhook/register", addCall(db))
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
// @Router /v1/market [get]
func getTickerHandler(storage storage.Market) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		coin := c.Param("coin")
		coin = strings.ToUpper(coin)
		if coin == "" {
			emptyPage(c)
			return
		}
		market := c.DefaultQuery("market", "dex")
		token := c.Query("token")
		result, err := storage.GetTicker(market, coin, token)
		if err != nil {
			ginutils.RenderError(c, http.StatusInternalServerError, err.Error())
		}
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
// @Router /v1/market [post]
func getTickersHandler(storage storage.Market) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		var md TickerRequest
		if err := c.BindJSON(&md); err != nil {
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}
		ginutils.RenderSuccess(c, []blockatlas.Ticker{})
	}
}

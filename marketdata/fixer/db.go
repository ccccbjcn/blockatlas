package fixer

import (
	"context"
	"github.com/trustwallet/backend/common"
	"github.com/trustwallet/backend/db"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func saveRates(rates *LatestRatesResponse) {
	saveToDb(*rates)
}

func saveToDb(rates LatestRatesResponse) {
	models := make([]mongo.WriteModel, 0)

	for c, rate := range rates.Rates {
		filter := bson.M{
			"currency": c,
		}

		update := bson.M{
			"rate":           common.Float64toPrecision(rate, 4),
			"rate_timestamp": rates.Timestamp,
			"updated_at":     time.Now(),
		}

		model := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(bson.D{{"$set", update}}).SetUpsert(true)
		models = append(models, model)
	}
	_, err := db.FiatRates.BulkWrite(context.TODO(), models)
	if err != nil {
		logger.Error("Error saveRates :", err)
	} else {
		logger.Info("Fiat rates updated OK")
	}

}

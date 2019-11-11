package fixer

import "time"

type LatestRatesResponse struct {
	Timestamp int                `json:"timestamp"`
	Rates     map[string]float64 `json:"rates"`
	UpdatedAt time.Time          `json:"updated_at"`
}

package usecases

import "github.com/braejan/go-bia-challenge/internal/domain/consumption/entity"

type ConsumptionUsecases interface {
	// GetDailyAccumulatedConsumptionByIDs returns the daily accumulated consumption for
	// a list of meter ids and a date range.
	GetDailyAccumulatedConsumptionByIDs(ids []int64, beginDate, endDate string) (consumptions entity.ConsumptionResponse, err error)
	// GetWeeklyAccumulatedConsumptionByIDs returns the weekly accumulated consumption for
	// a list of meter ids and a date range.
	GetWeeklyAccumulatedConsumptionByIDs(ids []int64, beginDate, endDate string) (consumptions entity.ConsumptionResponse, err error)
	// GetMonthlyAccumulatedConsumptionByIDs returns the monthly accumulated consumption for
	// a list of meter ids and a date range.
	GetMonthlyAccumulatedConsumptionByIDs(ids []int64, beginDate, endDate string) (consumptions entity.ConsumptionResponse, err error)
}

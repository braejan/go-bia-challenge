package repository

import "github.com/braejan/go-bia-challenge/internal/domain/consumption/entity"

type ConsumptionRepository interface {
	// GetConsumptionsByIDSGroupedDaily returns a list of consumptions grouped by day.
	GetConsumptionsByIDSGroupedDaily(ids []int64, beginDate, endDate string) (consumptions []*entity.Consumption, err error)
	// GetConsumptionsByIDSGroupedWeekly returns a list of consumptions grouped by week.
	GetConsumptionsByIDSGroupedWeekly(ids []int64, beginDate, endDate string) (consumptions []*entity.Consumption, err error)
	// GetConsumptionsByIDSGroupedMonthly returns a list of consumptions grouped by month.
	GetConsumptionsByIDSGroupedMonthly(ids []int64, beginDate, endDate string) (consumptions []*entity.Consumption, err error)
}

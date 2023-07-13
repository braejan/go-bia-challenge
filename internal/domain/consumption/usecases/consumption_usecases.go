package usecases

import (
	"github.com/braejan/go-bia-challenge/internal/domain/consumption/entity"
	"github.com/braejan/go-bia-challenge/internal/domain/consumption/repository"
	"github.com/braejan/go-bia-challenge/internal/domain/consumption/util"
	"github.com/braejan/go-bia-challenge/internal/valueobject/consumption"
)

type ConsumptionUsecasesImpl struct {
	consumptionRepo repository.ConsumptionRepository
}

// NewConsumptionUsecases creates a new instance of ConsumptionUsecases interface implementation.
func NewConsumptionUsecases(consumptionRepo repository.ConsumptionRepository) (consumptionUsecases ConsumptionUsecases, err error) {
	if consumptionRepo == nil {
		return nil, consumption.ErrNilConsumptionRepository
	}
	return &ConsumptionUsecasesImpl{
		consumptionRepo: consumptionRepo,
	}, nil
}

// ConsumptionUsecases interface implementation.

func (c *ConsumptionUsecasesImpl) GetDailyAccumulatedConsumptionByIDs(ids string, beginDate, endDate string) (consumptions entity.ConsumptionResponse, err error) {
	// Get consumptions from repository for daily accumulated consumption.
	intIDs, err := util.String2Int64Slice(ids)
	if err != nil {
		return consumptions, err
	}
	consumptionsRepo, err := c.consumptionRepo.GetConsumptionsByIDSGroupedDaily(intIDs, beginDate, endDate)
	if err != nil {
		return consumptions, err
	}
	consumptions = *util.Consumption2ConsumptionResponse(consumptionsRepo)
	return
}

func (c *ConsumptionUsecasesImpl) GetWeeklyAccumulatedConsumptionByIDs(ids string, beginDate, endDate string) (consumptions entity.ConsumptionResponse, err error) {
	// Get consumptions from repository for daily accumulated consumption.
	intIDs, err := util.String2Int64Slice(ids)
	if err != nil {
		return consumptions, err
	}
	consumptionsRepo, err := c.consumptionRepo.GetConsumptionsByIDSGroupedWeekly(intIDs, beginDate, endDate)
	if err != nil {
		return consumptions, err
	}
	consumptions = *util.Consumption2ConsumptionResponse(consumptionsRepo)
	return
}

func (c *ConsumptionUsecasesImpl) GetMonthlyAccumulatedConsumptionByIDs(ids string, beginDate, endDate string) (consumptions entity.ConsumptionResponse, err error) {
	// Get consumptions from repository for monthly accumulated consumption.
	intIDs, err := util.String2Int64Slice(ids)
	if err != nil {
		return consumptions, err
	}
	consumptionsRepo, err := c.consumptionRepo.GetConsumptionsByIDSGroupedMonthly(intIDs, beginDate, endDate)
	if err != nil {
		return consumptions, err
	}
	consumptions = *util.Consumption2ConsumptionResponse(consumptionsRepo)
	return
}

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

func (c *ConsumptionUsecasesImpl) GetDailyAccumulatedConsumptionByIDs(ids []int64, beginDate, endDate string) (consumptions entity.ConsumptionResponse, err error) {
	// Get consumptions from repository for daily accumulated consumption.
	consumptionsRepo, err := c.consumptionRepo.GetConsumptionsByIDSGroupedDaily(ids, beginDate, endDate)
	if err != nil {
		return consumptions, err
	}
	consumptions = *util.Consumption2ConsumptionResponse(consumptionsRepo)
	return
}

func (c *ConsumptionUsecasesImpl) GetWeeklyAccumulatedConsumptionByIDs(ids []int64, beginDate, endDate string) (consumptions entity.ConsumptionResponse, err error) {
	consumptions = entity.ConsumptionResponse{}
	// Get consumptions from repository for daily accumulated consumption.
	consumptionsRepo, err := c.consumptionRepo.GetConsumptionsByIDSGroupedWeekly(ids, beginDate, endDate)
	if err != nil {
		return consumptions, err
	}
	consumptions = *util.Consumption2ConsumptionResponse(consumptionsRepo)
	return
}

func (c *ConsumptionUsecasesImpl) GetMonthlyAccumulatedConsumptionByIDs(ids []int64, beginDate, endDate string) (consumptions entity.ConsumptionResponse, err error) {
	// Get consumptions from repository for monthly accumulated consumption.
	consumptionsRepo, err := c.consumptionRepo.GetConsumptionsByIDSGroupedMonthly(ids, beginDate, endDate)
	if err != nil {
		return consumptions, err
	}
	consumptions = *util.Consumption2ConsumptionResponse(consumptionsRepo)
	return
}

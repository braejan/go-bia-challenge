package usecases

import (
	"math"

	"github.com/braejan/go-bia-challenge/internal/domain/consumption/entity"
	"github.com/braejan/go-bia-challenge/internal/domain/consumption/repository"
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
	consumptions = entity.ConsumptionResponse{}
	// Get consumptions from repository for daily accumulated consumption.
	consumptionsRepo, err := c.consumptionRepo.GetConsumptionsByIDSGroupedDaily(ids, beginDate, endDate)
	if err != nil {
		return consumptions, err
	}
	// Convert consumptions from repository to consumptions for daily accumulated consumption.
	currentID := int64(math.MinInt64)
	dataGraphIndex := -1
	for _, consumptionRepo := range consumptionsRepo {
		if currentID < consumptionRepo.MeterID {
			dataGraphIndex++
			newDataGraph := &entity.ConsumptionDataGraph{
				MeterID: consumptionRepo.MeterID,
				Address: "Dirección Mock",
			}
			consumptions.DataGraph = append(consumptions.DataGraph, newDataGraph)
			currentID = consumptionRepo.MeterID
		}
		if dataGraphIndex == 0 {
			consumptions.Period = append(consumptions.Period, consumptionRepo.Date)

		}
		currentDataGraph := consumptions.DataGraph[dataGraphIndex]
		currentDataGraph.ActiveEnergy = append(currentDataGraph.ActiveEnergy, consumptionRepo.ActiveEnergy)
		currentDataGraph.ReactiveEnergy = append(currentDataGraph.ReactiveEnergy, consumptionRepo.ReactiveEnergy)
		currentDataGraph.CapacitiveReactive = append(currentDataGraph.CapacitiveReactive, consumptionRepo.CapacitiveReactive)
		currentDataGraph.Solar = append(currentDataGraph.Solar, consumptionRepo.Solar)
	}
	return consumptions, nil
}

func (c *ConsumptionUsecasesImpl) GetWeeklyAccumulatedConsumptionByIDs(ids []int64, beginDate, endDate string) (consumptions entity.ConsumptionResponse, err error) {
	return
}

func (c *ConsumptionUsecasesImpl) GetMonthlyAccumulatedConsumptionByIDs(ids []int64, beginDate, endDate string) (consumptions entity.ConsumptionResponse, err error) {
	return
}

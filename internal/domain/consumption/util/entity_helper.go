package util

import (
	"math"

	"github.com/braejan/go-bia-challenge/internal/domain/consumption/entity"
)

// Consumption2ConsumptionResponse converts a consumption to a consumption response.
func Consumption2ConsumptionResponse(consumptions []*entity.Consumption) (consumptionResponse *entity.ConsumptionResponse) {
	// Convert consumptions from repository to consumptions for daily accumulated consumption.
	currentID := int64(math.MinInt64)
	dataGraphIndex := -1
	consumptionResponse = &entity.ConsumptionResponse{}
	for _, consumption := range consumptions {
		if currentID < consumption.MeterID {
			dataGraphIndex++
			newDataGraph := &entity.ConsumptionDataGraph{
				MeterID: consumption.MeterID,
				Address: "Dirección Mock",
			}
			consumptionResponse.DataGraph = append(consumptionResponse.DataGraph, newDataGraph)
			currentID = consumption.MeterID
		}
		if dataGraphIndex == 0 {
			consumptionResponse.Period = append(consumptionResponse.Period, consumption.Date)

		}
		currentDataGraph := consumptionResponse.DataGraph[dataGraphIndex]
		currentDataGraph.ActiveEnergy = append(currentDataGraph.ActiveEnergy, consumption.ActiveEnergy)
		currentDataGraph.ReactiveEnergy = append(currentDataGraph.ReactiveEnergy, consumption.ReactiveEnergy)
		currentDataGraph.CapacitiveReactive = append(currentDataGraph.CapacitiveReactive, consumption.CapacitiveReactive)
		currentDataGraph.Solar = append(currentDataGraph.Solar, consumption.Solar)
	}
	return
}

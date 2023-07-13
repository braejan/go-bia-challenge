package util_test

import (
	"testing"

	"github.com/braejan/go-bia-challenge/internal/domain/consumption/entity"
	"github.com/braejan/go-bia-challenge/internal/domain/consumption/util"
	"github.com/stretchr/testify/assert"
)

// TestConsumption2ConsumptionResponse with empty consumptions.
func TestConsumption2ConsumptionResponse_EmptyConsumptions(t *testing.T) {
	// Given empty consumptions
	consumptions := []*entity.Consumption{}
	// When call Consumption2ConsumptionResponse
	consumptionResponse := util.Consumption2ConsumptionResponse(consumptions)
	// Then return an empty consumption response
	assert.Empty(t, consumptionResponse.DataGraph)
	assert.Empty(t, consumptionResponse.Period)
}

// TestConsumption2ConsumptionResponse with a consumption.
func TestConsumption2ConsumptionResponse_OneConsumption(t *testing.T) {
	// Given a consumption
	consumptions := getConsumptionArrayForTest()[0]
	// When call Consumption2ConsumptionResponse
	consumptionResponse := util.Consumption2ConsumptionResponse([]*entity.Consumption{consumptions})
	// Then return a consumption response
	assert.Equal(t, 1, len(consumptionResponse.DataGraph))
	assert.NotEmpty(t, consumptionResponse.Period)
}

// TestConsumption2ConsumptionResponse with a consumption list.
func TestConsumption2ConsumptionResponse_ConsumptionList(t *testing.T) {
	// Given a consumption list
	consumptions := getConsumptionArrayForTest()
	// When call Consumption2ConsumptionResponse
	consumptionResponse := util.Consumption2ConsumptionResponse(consumptions)
	// Then return a consumption response
	assert.Equal(t, 2, len(consumptionResponse.DataGraph))
	assert.NotEmpty(t, consumptionResponse.Period)
}

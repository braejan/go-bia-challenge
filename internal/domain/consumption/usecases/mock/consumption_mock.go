package mock

import (
	"github.com/braejan/go-bia-challenge/internal/domain/consumption/entity"
	"github.com/stretchr/testify/mock"
)

type mockConsumptionUsecases struct {
	mock.Mock
}

// NewMockConsumptionUsecases returns a mock of the consumption usecases.
func NewMockConsumptionUsecases() *mockConsumptionUsecases {
	return &mockConsumptionUsecases{}
}

// GetDailyAccumulatedConsumptionByIDs returns the daily accumulated consumption for
// a list of meter ids and a date range.
func (m *mockConsumptionUsecases) GetDailyAccumulatedConsumptionByIDs(ids string, beginDate, endDate string) (consumptions entity.ConsumptionResponse, err error) {
	args := m.Called(ids, beginDate, endDate)
	return args.Get(0).(entity.ConsumptionResponse), args.Error(1)
}

// GetWeeklyAccumulatedConsumptionByIDs returns the weekly accumulated consumption for
// a list of meter ids and a date range.
func (m *mockConsumptionUsecases) GetWeeklyAccumulatedConsumptionByIDs(ids string, beginDate, endDate string) (consumptions entity.ConsumptionResponse, err error) {
	args := m.Called(ids, beginDate, endDate)
	return args.Get(0).(entity.ConsumptionResponse), args.Error(1)
}

// GetMonthlyAccumulatedConsumptionByIDs returns the monthly accumulated consumption for
// a list of meter ids and a date range.
func (m *mockConsumptionUsecases) GetMonthlyAccumulatedConsumptionByIDs(ids string, beginDate, endDate string) (consumptions entity.ConsumptionResponse, err error) {
	args := m.Called(ids, beginDate, endDate)
	return args.Get(0).(entity.ConsumptionResponse), args.Error(1)
}

package mock

import (
	"github.com/braejan/go-bia-challenge/internal/domain/consumption/entity"
	"github.com/stretchr/testify/mock"
)

type mockConsumptionPostgresRepository struct {
	mock.Mock
}

// NewMockConsumptionPostgresRepository returns a mock of the consumption postgres repository.
func NewMockConsumptionPostgresRepository() *mockConsumptionPostgresRepository {
	return &mockConsumptionPostgresRepository{}
}

// GetConsumptionsByIDSGroupedDaily returns a list of consumptions grouped by day.
func (m *mockConsumptionPostgresRepository) GetConsumptionsByIDSGroupedDaily(ids []int64, beginDate, endDate string) (consumptions []*entity.Consumption, err error) {
	ret := m.Called(ids, beginDate, endDate)
	var r0 []*entity.Consumption
	if rf, ok := ret.Get(0).(func([]int64, string, string) []*entity.Consumption); ok {
		r0 = rf(ids, beginDate, endDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Consumption)
		}
	}
	var r1 error
	if rf, ok := ret.Get(1).(func([]int64, string, string) error); ok {
		r1 = rf(ids, beginDate, endDate)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// GetConsumptionsByIDSGroupedWeekly returns a list of consumptions grouped by week.
func (m *mockConsumptionPostgresRepository) GetConsumptionsByIDSGroupedWeekly(ids []int64, beginDate, endDate string) (consumptions []*entity.Consumption, err error) {
	ret := m.Called(ids, beginDate, endDate)
	var r0 []*entity.Consumption
	if rf, ok := ret.Get(0).(func([]int64, string, string) []*entity.Consumption); ok {
		r0 = rf(ids, beginDate, endDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Consumption)
		}
	}
	var r1 error
	if rf, ok := ret.Get(1).(func([]int64, string, string) error); ok {
		r1 = rf(ids, beginDate, endDate)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// GetConsumptionsByIDSGroupedMonthly returns a list of consumptions grouped by month.
func (m *mockConsumptionPostgresRepository) GetConsumptionsByIDSGroupedMonthly(ids []int64, beginDate, endDate string) (consumptions []*entity.Consumption, err error) {
	ret := m.Called(ids, beginDate, endDate)
	var r0 []*entity.Consumption
	if rf, ok := ret.Get(0).(func([]int64, string, string) []*entity.Consumption); ok {
		r0 = rf(ids, beginDate, endDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Consumption)
		}
	}
	var r1 error
	if rf, ok := ret.Get(1).(func([]int64, string, string) error); ok {
		r1 = rf(ids, beginDate, endDate)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

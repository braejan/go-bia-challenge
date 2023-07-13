package usecases_test

import (
	"errors"
	"testing"

	"github.com/braejan/go-bia-challenge/internal/domain/consumption/repository"
	"github.com/braejan/go-bia-challenge/internal/domain/consumption/repository/postgres/mock"
	"github.com/braejan/go-bia-challenge/internal/domain/consumption/usecases"
	consumptionVO "github.com/braejan/go-bia-challenge/internal/valueobject/consumption"
	"github.com/braejan/go-bia-challenge/tests/data"
	"github.com/stretchr/testify/assert"
)

// TestNewConsumptionUsecases with nil consumption repository.
func TestNewConsumptionUsecases(t *testing.T) {
	// Given a nil consumption repository
	var consumptionRepo repository.ConsumptionRepository
	// When call NewConsumptionUsecases
	_, err := usecases.NewConsumptionUsecases(consumptionRepo)
	// Then return an error
	assert.NotNil(t, err)
	assert.Equal(t, consumptionVO.ErrNilConsumptionRepository, err)
}

// TestNewConsumptionUsecases with a valid consumption repository.
func TestNewConsumptionUsecases_ValidConsumptionRepository(t *testing.T) {
	// Given a valid consumption repository
	consumptionRepo := mock.NewMockConsumptionPostgresRepository()
	// When call NewConsumptionUsecases
	consumptionUsecases, err := usecases.NewConsumptionUsecases(consumptionRepo)
	// Then return a consumption usecases
	assert.Nil(t, err)
	assert.NotNil(t, consumptionUsecases)
}

// TestGetDailyAccumulatedConsumptionByIDs with invalid ids.
func TestGetDailyAccumulatedConsumptionByIDs_InvalidIDs(t *testing.T) {
	// Given invalid ids
	ids := "a"
	beginDate := "2020-01-01"
	endDate := "2020-01-02"
	consumptionRepo := mock.NewMockConsumptionPostgresRepository()
	consumptionUsecases, _ := usecases.NewConsumptionUsecases(consumptionRepo)
	// When call GetDailyAccumulatedConsumptionByIDs
	consumptions, err := consumptionUsecases.GetDailyAccumulatedConsumptionByIDs(ids, beginDate, endDate)
	// Then return an error
	assert.NotNil(t, err)
	assert.Empty(t, consumptions)
}

// TestGetDailyAccumulatedConsumptionByIDs with invalid begin date.
func TestGetDailyAccumulatedConsumptionByIDs_InvalidBeginDate(t *testing.T) {
	// Given invalid begin date
	ids := "1,2,3"
	beginDate := "2020-01-0"
	endDate := "2020-01-02"
	consumptionRepo := mock.NewMockConsumptionPostgresRepository()
	consumptionUsecases, _ := usecases.NewConsumptionUsecases(consumptionRepo)
	// When call GetDailyAccumulatedConsumptionByIDs
	consumptions, err := consumptionUsecases.GetDailyAccumulatedConsumptionByIDs(ids, beginDate, endDate)
	// Then return an error
	assert.NotNil(t, err)
	assert.Empty(t, consumptions)
}

// TestGetDailyAccumulatedConsumptionByIDs with invalid end date.
func TestGetDailyAccumulatedConsumptionByIDs_InvalidEndDate(t *testing.T) {
	// Given invalid end date
	ids := "1,2,3"
	beginDate := "2020-01-01"
	endDate := "2020-01-0"
	consumptionRepo := mock.NewMockConsumptionPostgresRepository()
	consumptionUsecases, _ := usecases.NewConsumptionUsecases(consumptionRepo)
	// When call GetDailyAccumulatedConsumptionByIDs
	consumptions, err := consumptionUsecases.GetDailyAccumulatedConsumptionByIDs(ids, beginDate, endDate)
	// Then return an error
	assert.NotNil(t, err)
	assert.Empty(t, consumptions)
}

// TestGetDailyAccumulatedConsumptionByIDs with valid parameters and a mocked error
// calling GetConsumptionsByIDSGroupedDaily
func TestGetDailyAccumulatedConsumptionByIDs_ValidParameters_MockedError(t *testing.T) {
	// Given valid parameters and a mocked error calling GetConsumptionsByIDSGroupedDaily
	ids := "1,2,3"
	beginDate := "2020-01-01"
	endDate := "2020-01-02"
	consumptionRepo := mock.NewMockConsumptionPostgresRepository()
	consumptionRepo.On(
		"GetConsumptionsByIDSGroupedDaily",
		[]int64{1, 2, 3},
		"2020-01-01",
		"2020-01-02").Return(nil, errors.New("mocked error"))
	consumptionUsecases, _ := usecases.NewConsumptionUsecases(consumptionRepo)
	// When call GetDailyAccumulatedConsumptionByIDs
	consumptions, err := consumptionUsecases.GetDailyAccumulatedConsumptionByIDs(ids, beginDate, endDate)
	// Then return an error
	assert.NotNil(t, err)
	assert.Empty(t, consumptions)
}

// TestGetDailyAccumulatedConsumptionByIDs with valid parameters and a mocked response
// calling GetConsumptionsByIDSGroupedDaily
func TestGetDailyAccumulatedConsumptionByIDs_ValidParameters_MockedResponse(t *testing.T) {
	// Given valid parameters and a mocked response calling GetConsumptionsByIDSGroupedDaily
	ids := "1,2,3"
	beginDate := "2020-01-01"
	endDate := "2020-01-02"
	consumptionRepo := mock.NewMockConsumptionPostgresRepository()
	consumptionRepo.On(
		"GetConsumptionsByIDSGroupedDaily",
		[]int64{1, 2, 3},
		"2020-01-01",
		"2020-01-02").Return(data.GetConsumptionArrayForTest(), nil)
	consumptionUsecases, _ := usecases.NewConsumptionUsecases(consumptionRepo)
	// When call GetDailyAccumulatedConsumptionByIDs
	consumptions, err := consumptionUsecases.GetDailyAccumulatedConsumptionByIDs(ids, beginDate, endDate)
	// Then return a consumption response
	assert.Nil(t, err)
	assert.Equal(t, 2, len(consumptions.DataGraph))
	assert.NotEmpty(t, consumptions.Period)
}

// TestGetWeeklyAccumulatedConsumptionByIDs with invalid ids.
func TestGetWeeklyAccumulatedConsumptionByIDs_InvalidIDs(t *testing.T) {
	// Given invalid ids
	ids := "a"
	beginDate := "2020-01-01"
	endDate := "2020-01-02"
	consumptionRepo := mock.NewMockConsumptionPostgresRepository()
	consumptionUsecases, _ := usecases.NewConsumptionUsecases(consumptionRepo)
	// When call GetWeeklyAccumulatedConsumptionByIDs
	consumptions, err := consumptionUsecases.GetWeeklyAccumulatedConsumptionByIDs(ids, beginDate, endDate)
	// Then return an error
	assert.NotNil(t, err)
	assert.Empty(t, consumptions)
}

// TestGetWeeklyAccumulatedConsumptionByIDs with invalid begin date.
func TestGetWeeklyAccumulatedConsumptionByIDs_InvalidBeginDate(t *testing.T) {
	// Given invalid begin date
	ids := "1,2,3"
	beginDate := "2020-01-0"
	endDate := "2020-01-02"
	consumptionRepo := mock.NewMockConsumptionPostgresRepository()
	consumptionUsecases, _ := usecases.NewConsumptionUsecases(consumptionRepo)
	// When call GetWeeklyAccumulatedConsumptionByIDs
	consumptions, err := consumptionUsecases.GetWeeklyAccumulatedConsumptionByIDs(ids, beginDate, endDate)
	// Then return an error
	assert.NotNil(t, err)
	assert.Empty(t, consumptions)
}

// TestGetWeeklyAccumulatedConsumptionByIDs with invalid end date.
func TestGetWeeklyAccumulatedConsumptionByIDs_InvalidEndDate(t *testing.T) {
	// Given invalid end date
	ids := "1,2,3"
	beginDate := "2020-01-01"
	endDate := "2020-01-0"
	consumptionRepo := mock.NewMockConsumptionPostgresRepository()
	consumptionUsecases, _ := usecases.NewConsumptionUsecases(consumptionRepo)
	// When call GetWeeklyAccumulatedConsumptionByIDs
	consumptions, err := consumptionUsecases.GetWeeklyAccumulatedConsumptionByIDs(ids, beginDate, endDate)
	// Then return an error
	assert.NotNil(t, err)
	assert.Empty(t, consumptions)
}

// TestGetWeeklyAccumulatedConsumptionByIDs with valid parameters and a mocked error
// calling GetConsumptionsByIDSGroupedWeekly
func TestGetWeeklyAccumulatedConsumptionByIDs_ValidParameters_MockedError(t *testing.T) {
	// Given valid parameters and a mocked error calling GetConsumptionsByIDSGroupedWeekly
	ids := "1,2,3"
	beginDate := "2020-01-01"
	endDate := "2020-01-02"
	consumptionRepo := mock.NewMockConsumptionPostgresRepository()
	consumptionRepo.On(
		"GetConsumptionsByIDSGroupedWeekly",
		[]int64{1, 2, 3},
		"2020-01-01",
		"2020-01-02").Return(nil, errors.New("mocked error"))
	consumptionUsecases, _ := usecases.NewConsumptionUsecases(consumptionRepo)
	// When call GetWeeklyAccumulatedConsumptionByIDs
	consumptions, err := consumptionUsecases.GetWeeklyAccumulatedConsumptionByIDs(ids, beginDate, endDate)
	// Then return an error
	assert.NotNil(t, err)
	assert.Empty(t, consumptions)
}

// TestGetWeeklyAccumulatedConsumptionByIDs with valid parameters and a mocked response
// calling GetConsumptionsByIDSGroupedWeekly
func TestGetWeeklyAccumulatedConsumptionByIDs_ValidParameters_MockedResponse(t *testing.T) {
	// Given valid parameters and a mocked response calling GetConsumptionsByIDSGroupedWeekly
	ids := "1,2,3"
	beginDate := "2020-01-01"
	endDate := "2020-01-02"
	consumptionRepo := mock.NewMockConsumptionPostgresRepository()
	consumptionRepo.On(
		"GetConsumptionsByIDSGroupedWeekly",
		[]int64{1, 2, 3},
		"2020-01-01",
		"2020-01-02").Return(data.GetConsumptionArrayForTest(), nil)
	consumptionUsecases, _ := usecases.NewConsumptionUsecases(consumptionRepo)
	// When call GetWeeklyAccumulatedConsumptionByIDs
	consumptions, err := consumptionUsecases.GetWeeklyAccumulatedConsumptionByIDs(ids, beginDate, endDate)
	// Then return a consumption response
	assert.Nil(t, err)
	assert.Equal(t, 2, len(consumptions.DataGraph))
	assert.NotEmpty(t, consumptions.Period)
}

// TestGetMonthlyAccumulatedConsumptionByIDs with invalid ids.
func TestGetMonthlyAccumulatedConsumptionByIDs_InvalidIDs(t *testing.T) {
	// Given invalid ids
	ids := "a"
	beginDate := "2020-01-01"
	endDate := "2020-01-02"
	consumptionRepo := mock.NewMockConsumptionPostgresRepository()
	consumptionUsecases, _ := usecases.NewConsumptionUsecases(consumptionRepo)
	// When call GetMonthlyAccumulatedConsumptionByIDs
	consumptions, err := consumptionUsecases.GetMonthlyAccumulatedConsumptionByIDs(ids, beginDate, endDate)
	// Then return an error
	assert.NotNil(t, err)
	assert.Empty(t, consumptions)
}

// TestGetMonthlyAccumulatedConsumptionByIDs with invalid begin date.
func TestGetMonthlyAccumulatedConsumptionByIDs_InvalidBeginDate(t *testing.T) {
	// Given invalid begin date
	ids := "1,2,3"
	beginDate := "2020-01-0"
	endDate := "2020-01-02"
	consumptionRepo := mock.NewMockConsumptionPostgresRepository()
	consumptionUsecases, _ := usecases.NewConsumptionUsecases(consumptionRepo)
	// When call GetMonthlyAccumulatedConsumptionByIDs
	consumptions, err := consumptionUsecases.GetMonthlyAccumulatedConsumptionByIDs(ids, beginDate, endDate)
	// Then return an error
	assert.NotNil(t, err)
	assert.Empty(t, consumptions)
}

// TestGetMonthlyAccumulatedConsumptionByIDs with invalid end date.
func TestGetMonthlyAccumulatedConsumptionByIDs_InvalidEndDate(t *testing.T) {
	// Given invalid end date
	ids := "1,2,3"
	beginDate := "2020-01-01"
	endDate := "2020-01-0"
	consumptionRepo := mock.NewMockConsumptionPostgresRepository()
	consumptionUsecases, _ := usecases.NewConsumptionUsecases(consumptionRepo)
	// When call GetMonthlyAccumulatedConsumptionByIDs
	consumptions, err := consumptionUsecases.GetMonthlyAccumulatedConsumptionByIDs(ids, beginDate, endDate)
	// Then return an error
	assert.NotNil(t, err)
	assert.Empty(t, consumptions)
}

// TestGetMonthlyAccumulatedConsumptionByIDs with valid parameters and a mocked error
// calling GetConsumptionsByIDSGroupedMonthly
func TestGetMonthlyAccumulatedConsumptionByIDs_ValidParameters_MockedError(t *testing.T) {
	// Given valid parameters and a mocked error calling GetConsumptionsByIDSGroupedMonthly
	ids := "1,2,3"
	beginDate := "2020-01-01"
	endDate := "2020-01-02"
	consumptionRepo := mock.NewMockConsumptionPostgresRepository()
	consumptionRepo.On(
		"GetConsumptionsByIDSGroupedMonthly",
		[]int64{1, 2, 3},
		"2020-01-01",
		"2020-01-02").Return(nil, errors.New("mocked error"))
	consumptionUsecases, _ := usecases.NewConsumptionUsecases(consumptionRepo)
	// When call GetMonthlyAccumulatedConsumptionByIDs
	consumptions, err := consumptionUsecases.GetMonthlyAccumulatedConsumptionByIDs(ids, beginDate, endDate)
	// Then return an error
	assert.NotNil(t, err)
	assert.Empty(t, consumptions)
}

// TestGetMonthlyAccumulatedConsumptionByIDs with valid parameters and a mocked response
// calling GetConsumptionsByIDSGroupedMonthly
func TestGetMonthlyAccumulatedConsumptionByIDs_ValidParameters_MockedResponse(t *testing.T) {
	// Given valid parameters and a mocked response calling GetConsumptionsByIDSGroupedMonthly
	ids := "1,2,3"
	beginDate := "2020-01-01"
	endDate := "2020-01-02"
	consumptionRepo := mock.NewMockConsumptionPostgresRepository()
	consumptionRepo.On(
		"GetConsumptionsByIDSGroupedMonthly",
		[]int64{1, 2, 3},
		"2020-01-01",
		"2020-01-02").Return(data.GetConsumptionArrayForTest(), nil)
	consumptionUsecases, _ := usecases.NewConsumptionUsecases(consumptionRepo)
	// When call GetMonthlyAccumulatedConsumptionByIDs
	consumptions, err := consumptionUsecases.GetMonthlyAccumulatedConsumptionByIDs(ids, beginDate, endDate)
	// Then return a consumption response
	assert.Nil(t, err)
	assert.Equal(t, 2, len(consumptions.DataGraph))
	assert.NotEmpty(t, consumptions.Period)
}

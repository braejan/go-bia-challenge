package postgres_test

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/braejan/go-bia-challenge/internal/domain/consumption/repository/postgres"
	mockvoPostgres "github.com/braejan/go-bia-challenge/internal/valueobject/postgres/mock"
	"github.com/braejan/go-bia-challenge/tests/data"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

// TestGetConsumptionsByIDSGroupedDaily with error opening database.
func TestGetConsumptionsByIDSGroupedDaily_ErrorOpeningDatabase(t *testing.T) {
	// Given a mocked PostgresDatabase with error opening database
	mockPostgresDatabase := mockvoPostgres.NewMockBasePostgresDatabase()
	mockPostgresDatabase.On("Open").Return(nil, errors.New("mocked error"))
	// And a valid repository
	postgresRepo := postgres.NewConsumptionPostgresRepository(mockPostgresDatabase)
	// When call GetConsumptionsByIDSGroupedDaily
	_, err := postgresRepo.GetConsumptionsByIDSGroupedDaily([]int64{1, 2}, "2021-01-01", "2021-01-02")
	// Then return an error
	assert.NotNil(t, err)
}

// TestGetConsumptionsByIDSGroupedDaily with error querying database.
func TestGetConsumptionsByIDSGroupedDaily_ErrorQueryingDatabase(t *testing.T) {
	// Given a mocked PostgresDatabase with error querying database
	mockPostgresDatabase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a sqlmock valid database
	dbMocked, dbSqlmock, _ := sqlmock.New()
	defer dbMocked.Close()
	mockPostgresDatabase.On("Open").Return(dbMocked, nil)
	mockPostgresDatabase.On("Close", dbMocked).Return(nil)
	dbSqlmock.ExpectQuery("SELECT (.+) FROM consumptions WHERE (.+) GROUP BY (.+)").WillReturnError(errors.New("mocked error"))
	// And a valid repository
	postgresRepo := postgres.NewConsumptionPostgresRepository(mockPostgresDatabase)
	// When call GetConsumptionsByIDSGroupedDaily
	_, err := postgresRepo.GetConsumptionsByIDSGroupedDaily([]int64{1, 2}, "2021-01-01", "2021-01-02")
	// Then return an error
	assert.NotNil(t, err)
}

// TestGetConsumptionsByIDSGroupedDaily with valid consumptions.
func TestGetConsumptionsByIDSGroupedDaily_ValidConsumptions(t *testing.T) {
	// Given a mocked PostgresDatabase with error querying database
	mockPostgresDatabase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a sqlmock valid database
	dbMocked, dbSqlmock, _ := sqlmock.New()
	defer dbMocked.Close()
	mockPostgresDatabase.On("Open").Return(dbMocked, nil)
	mockPostgresDatabase.On("Close", dbMocked).Return(nil)
	metersID := []int64{1, 2}
	beginDate := "2023-06-01"
	endDate := "2023-07-10"
	consumption := data.GetConsumptionArrayForTest()[0]
	expected := sqlmock.NewRows([]string{"day", "date_string", "meter_id", "total_active_energy", "total_reactive_energy", "total_capacitive_reactive", "total_solar"})
	expected.AddRow(time.Time{}, consumption.Date, consumption.MeterID, consumption.ActiveEnergy, consumption.ReactiveEnergy, consumption.CapacitiveReactive, consumption.Solar)
	dbSqlmock.ExpectQuery("SELECT (.+) FROM consumptions WHERE (.+) GROUP BY (.+)").WithArgs(pq.Array(metersID), beginDate, endDate).WillReturnRows(expected)
	// And a valid repository
	postgresRepo := postgres.NewConsumptionPostgresRepository(mockPostgresDatabase)
	// When call GetConsumptionsByIDSGroupedDaily
	consumptions, err := postgresRepo.GetConsumptionsByIDSGroupedDaily(metersID, beginDate, endDate)
	// Then return a valid consumption
	assert.Nil(t, err)
	assert.Equal(t, 1, len(consumptions))
}

// TestGetConsumptionsByIDSGroupedWeekly with error opening database.
func TestGetConsumptionsByIDSGroupedWeekly_ErrorOpeningDatabase(t *testing.T) {
	// Given a mocked PostgresDatabase with error opening database
	mockPostgresDatabase := mockvoPostgres.NewMockBasePostgresDatabase()
	mockPostgresDatabase.On("Open").Return(nil, errors.New("mocked error"))
	// And a valid repository
	postgresRepo := postgres.NewConsumptionPostgresRepository(mockPostgresDatabase)
	// When call GetConsumptionsByIDSGroupedWeekly
	_, err := postgresRepo.GetConsumptionsByIDSGroupedWeekly([]int64{1, 2}, "2021-01-01", "2021-01-02")
	// Then return an error
	assert.NotNil(t, err)
}

// TestGetConsumptionsByIDSGroupedWeekly with error querying database.
func TestGetConsumptionsByIDSGroupedWeekly_ErrorQueryingDatabase(t *testing.T) {
	// Given a mocked PostgresDatabase with error querying database
	mockPostgresDatabase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a sqlmock valid database
	dbMocked, dbSqlmock, _ := sqlmock.New()
	defer dbMocked.Close()
	mockPostgresDatabase.On("Open").Return(dbMocked, nil)
	mockPostgresDatabase.On("Close", dbMocked).Return(nil)
	dbSqlmock.ExpectQuery("SELECT (.+) FROM consumptions WHERE (.+) GROUP BY (.+)").WillReturnError(errors.New("mocked error"))
	// And a valid repository
	postgresRepo := postgres.NewConsumptionPostgresRepository(mockPostgresDatabase)
	// When call GetConsumptionsByIDSGroupedWeekly
	_, err := postgresRepo.GetConsumptionsByIDSGroupedWeekly([]int64{1, 2}, "2021-01-01", "2021-01-02")
	// Then return an error
	assert.NotNil(t, err)
}

// TestGetConsumptionsByIDSGroupedWeekly with valid consumptions.
func TestGetConsumptionsByIDSGroupedWeekly_ValidConsumptions(t *testing.T) {
	// Given a mocked PostgresDatabase with error querying database
	mockPostgresDatabase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a sqlmock valid database
	dbMocked, dbSqlmock, _ := sqlmock.New()
	defer dbMocked.Close()
	mockPostgresDatabase.On("Open").Return(dbMocked, nil)
	mockPostgresDatabase.On("Close", dbMocked).Return(nil)
	metersID := []int64{1, 2}
	beginDate := "2023-06-01"
	endDate := "2023-07-10"
	consumption := data.GetConsumptionArrayForTest()[0]
	expected := sqlmock.NewRows([]string{"weekly_period", "weekly_date_range", "meter_id", "total_active_energy", "total_reactive_energy", "total_capacitive_reactive", "total_solar"})
	expected.AddRow(time.Time{}, consumption.Date, consumption.MeterID, consumption.ActiveEnergy, consumption.ReactiveEnergy, consumption.CapacitiveReactive, consumption.Solar)
	dbSqlmock.ExpectQuery("SELECT (.+) FROM consumptions c WHERE (.+) GROUP BY (.+)").WithArgs(pq.Array(metersID), beginDate, endDate).WillReturnRows(expected)
	// And a valid repository
	postgresRepo := postgres.NewConsumptionPostgresRepository(mockPostgresDatabase)
	// When call GetConsumptionsByIDSGroupedWeekly
	consumptions, err := postgresRepo.GetConsumptionsByIDSGroupedWeekly(metersID, beginDate, endDate)
	// Then return a valid consumption
	assert.Nil(t, err)
	assert.Equal(t, 1, len(consumptions))
}

// TestGetConsumptionsByIDSGroupedMonthly with error opening database.
func TestGetConsumptionsByIDSGroupedMonthly_ErrorOpeningDatabase(t *testing.T) {
	// Given a mocked PostgresDatabase with error opening database
	mockPostgresDatabase := mockvoPostgres.NewMockBasePostgresDatabase()
	mockPostgresDatabase.On("Open").Return(nil, errors.New("mocked error"))
	// And a valid repository
	postgresRepo := postgres.NewConsumptionPostgresRepository(mockPostgresDatabase)
	// When call GetConsumptionsByIDSGroupedMonthly
	_, err := postgresRepo.GetConsumptionsByIDSGroupedMonthly([]int64{1, 2}, "2021-01-01", "2021-01-02")
	// Then return an error
	assert.NotNil(t, err)
}

// TestGetConsumptionsByIDSGroupedMonthly with error querying database.
func TestGetConsumptionsByIDSGroupedMonthly_ErrorQueryingDatabase(t *testing.T) {
	// Given a mocked PostgresDatabase with error querying database
	mockPostgresDatabase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a sqlmock valid database
	dbMocked, dbSqlmock, _ := sqlmock.New()
	defer dbMocked.Close()
	mockPostgresDatabase.On("Open").Return(dbMocked, nil)
	mockPostgresDatabase.On("Close", dbMocked).Return(nil)
	dbSqlmock.ExpectQuery("SELECT (.+) FROM consumptions WHERE (.+) GROUP BY (.+)").WillReturnError(errors.New("mocked error"))
	// And a valid repository
	postgresRepo := postgres.NewConsumptionPostgresRepository(mockPostgresDatabase)
	// When call GetConsumptionsByIDSGroupedMonthly
	_, err := postgresRepo.GetConsumptionsByIDSGroupedMonthly([]int64{1, 2}, "2021-01-01", "2021-01-02")
	// Then return an error
	assert.NotNil(t, err)
}

// TestGetConsumptionsByIDSGroupedMonthly with valid consumptions.
func TestGetConsumptionsByIDSGroupedMonthly_ValidConsumptions(t *testing.T) {
	// Given a mocked PostgresDatabase with error querying database
	mockPostgresDatabase := mockvoPostgres.NewMockBasePostgresDatabase()
	// And a sqlmock valid database
	dbMocked, dbSqlmock, _ := sqlmock.New()
	defer dbMocked.Close()
	mockPostgresDatabase.On("Open").Return(dbMocked, nil)
	mockPostgresDatabase.On("Close", dbMocked).Return(nil)
	metersID := []int64{1, 2}
	beginDate := "2023-06-01"
	endDate := "2023-07-10"
	consumption := data.GetConsumptionArrayForTest()[0]
	expected := sqlmock.NewRows([]string{"monthly_period", "monthly_date_period", "meter_id", "total_active_energy", "total_reactive_energy", "total_capacitive_reactive", "total_solar"})
	expected.AddRow(time.Time{}, consumption.Date, consumption.MeterID, consumption.ActiveEnergy, consumption.ReactiveEnergy, consumption.CapacitiveReactive, consumption.Solar)
	dbSqlmock.ExpectQuery("SELECT (.+) FROM consumptions c WHERE (.+) GROUP BY (.+)").WithArgs(pq.Array(metersID), beginDate, endDate).WillReturnRows(expected)
	// And a valid repository
	postgresRepo := postgres.NewConsumptionPostgresRepository(mockPostgresDatabase)
	// When call GetConsumptionsByIDSGroupedWeekly
	consumptions, err := postgresRepo.GetConsumptionsByIDSGroupedMonthly(metersID, beginDate, endDate)
	// Then return a valid consumption
	assert.Nil(t, err)
	assert.Equal(t, 1, len(consumptions))
}

package postgres

import (
	"context"

	"github.com/braejan/go-bia-challenge/internal/domain/consumption/entity"
	"github.com/braejan/go-bia-challenge/internal/valueobject/postgres"
	"github.com/lib/pq"
)

type ConsumptionPostgresRepository struct {
	baseDB postgres.PostgresDatabase
}

// NewConsumptionPostgresRepository creates a new instance of ConsumptionRepository interface implementation.
func NewConsumptionPostgresRepository(baseDB postgres.PostgresDatabase) *ConsumptionPostgresRepository {
	return &ConsumptionPostgresRepository{
		baseDB: baseDB,
	}
}

// ConsumptionRepository interface implementation.

// GetConsumptionsByIDSGroupedDaily returns a list of consumptions grouped by day.
const (
	getConsumptionsByIDSGroupedDailyQuery = `
	SELECT
    date::date AS day,
    TO_CHAR(date::date, 'MON DD') AS date_string,
    meter_id,
    SUM(active_energy) AS total_active_energy,
    SUM(reactive_energy) AS total_reactive_energy,
    SUM(capacitive_reactive) AS total_capacitive_reactive,
    SUM(solar) AS total_solar
FROM
    consumptions
WHERE
    meter_id = ANY ($1)
    AND date >= $2
    AND date <= $3
GROUP BY
    date::date,
    day,
    meter_id
ORDER by
	meter_id,
    day;
`
)

func (postgresRepo *ConsumptionPostgresRepository) GetConsumptionsByIDSGroupedDaily(ids []int64, beginDate, endDate string) (consumptions []*entity.Consumption, err error) {
	db, err := postgresRepo.baseDB.Open()
	if err != nil {
		return nil, err
	}
	defer postgresRepo.baseDB.Close(db)
	idArray := pq.Array(ids)
	rows, err := db.QueryContext(context.Background(), getConsumptionsByIDSGroupedDailyQuery, idArray, beginDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var consumption entity.Consumption
		var dailyPeriod string
		err = rows.Scan(
			&dailyPeriod,
			&consumption.Date,
			&consumption.MeterID,
			&consumption.ActiveEnergy,
			&consumption.ReactiveEnergy,
			&consumption.CapacitiveReactive,
			&consumption.Solar,
		)
		if err != nil {
			return nil, err
		}
		consumptions = append(consumptions, &consumption)
	}
	return
}

// GetConsumptionsByIDSGroupedWeekly returns a list of consumptions grouped by week.
const (
	getConsumptionsByIDSGroupedWeeklyQuery = `
	SELECT
    DATE_TRUNC('week', date) AS weekly_period,
    CONCAT(TO_CHAR(MIN(date), 'MON DD'), ' - ', TO_CHAR(MAX(date), 'MON DD')) AS weekly_date_range,
    meter_id,
    SUM(active_energy) AS total_active_energy,
    SUM(reactive_energy) AS total_reactive_energy,
    SUM(capacitive_reactive) AS total_capacitive_reactive,
    SUM(solar) AS total_solar
FROM
    consumptions c
WHERE
    meter_id = ANY ($1)
    AND date >= $2
    AND date <= $3
	GROUP BY
    weekly_period,
    meter_id
ORDER by
	meter_id asc,
    weekly_period desc
`
)

func (postgresRepo *ConsumptionPostgresRepository) GetConsumptionsByIDSGroupedWeekly(ids []int64, beginDate, endDate string) (consumptions []*entity.Consumption, err error) {
	db, err := postgresRepo.baseDB.Open()
	if err != nil {
		return nil, err
	}
	defer postgresRepo.baseDB.Close(db)
	idArray := pq.Array(ids)
	rows, err := db.QueryContext(context.Background(), getConsumptionsByIDSGroupedWeeklyQuery, idArray, beginDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var consumption entity.Consumption
		var weeklyPeriod string
		err = rows.Scan(
			&weeklyPeriod,
			&consumption.Date,
			&consumption.MeterID,
			&consumption.ActiveEnergy,
			&consumption.ReactiveEnergy,
			&consumption.CapacitiveReactive,
			&consumption.Solar,
		)
		if err != nil {
			return nil, err
		}
		consumptions = append(consumptions, &consumption)
	}
	return
}

// GetConsumptionsByIDSGroupedMonthly returns a list of consumptions grouped by month.
const (
	getConsumptionsByIDSGroupedMonthlyQuery = `
	SELECT
	DATE_TRUNC('month', date) AS monthly_period,
	TO_CHAR(MIN(date), 'MON YYYY') AS monthly_date_period,
	meter_id,
	SUM(active_energy) AS total_active_energy,
	SUM(reactive_energy) AS total_reactive_energy,
	SUM(capacitive_reactive) AS total_capacitive_reactive,
	SUM(solar) AS total_solar
FROM
	consumptions c
WHERE
	meter_id = ANY ($1)
	AND date >= $2
	AND date <= $3
GROUP BY
	monthly_period,
	meter_id
ORDER by
	meter_id asc,
	monthly_period asc
`
)

func (postgresRepo *ConsumptionPostgresRepository) GetConsumptionsByIDSGroupedMonthly(ids []int64, beginDate, endDate string) (consumptions []*entity.Consumption, err error) {
	db, err := postgresRepo.baseDB.Open()
	if err != nil {
		return nil, err
	}
	defer postgresRepo.baseDB.Close(db)
	idArray := pq.Array(ids)
	rows, err := db.QueryContext(context.Background(), getConsumptionsByIDSGroupedMonthlyQuery, idArray, beginDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var consumption entity.Consumption
		var monthlyPeriod string
		err = rows.Scan(
			&monthlyPeriod,
			&consumption.Date,
			&consumption.MeterID,
			&consumption.ActiveEnergy,
			&consumption.ReactiveEnergy,
			&consumption.CapacitiveReactive,
			&consumption.Solar,
		)
		if err != nil {
			return nil, err
		}
		consumptions = append(consumptions, &consumption)
	}
	return
}
